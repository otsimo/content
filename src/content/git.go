package content

type GitClient struct {
	Path   string
	GitUrl string
}

func (gc GitClient) Clone() {
	
}

func (gc GitClient) Fetch() {

}

func (gc GitClient) Pull() {

}

func NewGitClient(folder, url string) *GitClient {
	return &GitClient{Path: folder, GitUrl: url}
}
