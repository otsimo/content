// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!


export interface ABTestGroup {
  name?: string;
  appId?: string;
  weight?: number;
  active?: boolean;
}

export interface ABDisableReq {
  name?: string;
  appId?: string;
}

export interface ListTestGroupReq {
  appId?: string;
}

export interface AppTestingStatus {
  version?: number;
  tests?: ABTestGroup[];
}

export interface CheckUserStatusReq {
  userId?: string;
}

export interface ChangeUserTestGroupReq {
  userId?: string;
  appId?: string;
/**
TestGroupName is empty for no testing group
*/
  testGroupName?: string;
}

export interface UserTestGroupStatus {
  userId?: string;
  apps?: UserTestGroupStatusAppGroup[];
}

export interface UserTestGroupStatusAppGroup {
  appId?: string;
  testGroup?: string;
  version?: number;
}
