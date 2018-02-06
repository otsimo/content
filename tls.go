package content

// Copyright 2016 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"crypto/tls"
	"crypto/x509"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
	"google.golang.org/grpc/credentials"
)

type CertificateManager struct {
	sync.RWMutex
	certFile    string
	keyFile     string
	certificate *tls.Certificate
	x509Cert    *x509.Certificate

	Error   chan error
	watcher *fsnotify.Watcher
}

func GrpcCredentials(certFile, keyFile string) (credentials.TransportCredentials, error) {
	cm, err := NewCertificateManager(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	cfg := &tls.Config{
		GetCertificate: cm.GetCertificate,
	}
	return credentials.NewTLS(cfg), nil
}

func NewCertificateManager(certFile, keyFile string) (*CertificateManager, error) {
	cm := &CertificateManager{
		certFile: certFile,
		keyFile:  keyFile,
		Error:    make(chan error, 10),
	}
	err := cm.loadCertificate()
	if err != nil {
		logrus.WithError(err).Errorf("failed to load loadCertificate")
		return nil, err
	}
	go cm.watchCertificate()
	return cm, nil
}

func (cm *CertificateManager) GetCertificate(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	cm.RLock()
	defer cm.RUnlock()
	return cm.certificate, nil
}

func (cm *CertificateManager) loadCertificate() error {
	logrus.Println("Loading TLS certificates...")
	c, err := tls.LoadX509KeyPair(cm.certFile, cm.keyFile)
	if err != nil {
		return err
	}
	x509Cert, err := x509.ParseCertificate(c.Certificate[0])
	if err != nil {
		return err
	}
	cm.Lock()
	defer cm.Unlock()
	if cm.x509Cert != nil {
		if cm.x509Cert.NotBefore.Before(x509Cert.NotBefore) {
			return nil
		}
	}
	cm.certificate = &c
	cm.x509Cert = x509Cert
	return nil
}

func (cm *CertificateManager) watchCertificate() error {
	logrus.Println("Watching for TLS certificate changes...")
	err := cm.newWatcher()
	if err != nil {
		logrus.WithError(err).Errorf("failed to create file watcher")
		return err
	}
	for {
		select {
		case event := <-cm.watcher.Events:
			logrus.Infof("Receive new event %s", event.String())
			logrus.Println("Reloading TLS certificates...")
			err := cm.loadCertificate()
			if err != nil {
				cm.Error <- err
			}
			logrus.Println("Reloading TLS certificates complete.")
			err = cm.resetWatcher()
			if err != nil {
				cm.Error <- err
			}
		case err := <-cm.watcher.Errors:
			logrus.Infof("Receive new error %v", err)
			cm.Error <- err
		}
	}
}

func (cm *CertificateManager) newWatcher() error {
	var err error
	cm.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	err = cm.watcher.Add(cm.certFile)
	if err != nil {
		return err
	}
	return cm.watcher.Add(cm.keyFile)
}

func (cm *CertificateManager) resetWatcher() error {
	err := cm.watcher.Close()
	if err != nil {
		return err
	}
	return cm.newWatcher()
}
