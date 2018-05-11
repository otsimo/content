// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!

import * as apipb_models from "./models.pb";

export type GameEntryRequestRequestType =  "ADD"  | "ACTIVATE"  | "DEACTIVATE"  | "SETTINGS"  | "INDEX"  | "LOCALSETTINGS" ;
export const GameEntryRequestRequestType_ADD: GameEntryRequestRequestType = "ADD";
export const GameEntryRequestRequestType_ACTIVATE: GameEntryRequestRequestType = "ACTIVATE";
export const GameEntryRequestRequestType_DEACTIVATE: GameEntryRequestRequestType = "DEACTIVATE";
export const GameEntryRequestRequestType_SETTINGS: GameEntryRequestRequestType = "SETTINGS";
export const GameEntryRequestRequestType_INDEX: GameEntryRequestRequestType = "INDEX";
export const GameEntryRequestRequestType_LOCALSETTINGS: GameEntryRequestRequestType = "LOCALSETTINGS";

export type ListGamesRequestInnerState =  "ANY"  | "CREATED"  | "DEVELOPMENT"  | "WAITING"  | "REJECTED"  | "VALIDATED"  | "PRODUCTION" ;
export const ListGamesRequestInnerState_ANY: ListGamesRequestInnerState = "ANY";
export const ListGamesRequestInnerState_CREATED: ListGamesRequestInnerState = "CREATED";
export const ListGamesRequestInnerState_DEVELOPMENT: ListGamesRequestInnerState = "DEVELOPMENT";
export const ListGamesRequestInnerState_WAITING: ListGamesRequestInnerState = "WAITING";
export const ListGamesRequestInnerState_REJECTED: ListGamesRequestInnerState = "REJECTED";
export const ListGamesRequestInnerState_VALIDATED: ListGamesRequestInnerState = "VALIDATED";
export const ListGamesRequestInnerState_PRODUCTION: ListGamesRequestInnerState = "PRODUCTION";

export type RequestReleaseState =  "PRODUCTION_STATE"  | "ALL_STATES" ;
export const RequestReleaseState_PRODUCTION_STATE: RequestReleaseState = "PRODUCTION_STATE";
export const RequestReleaseState_ALL_STATES: RequestReleaseState = "ALL_STATES";

/**
A label selector operator is the set of operators that can be used in
a label selector requirement.
*/
export type LabelSelectorOperator =  "In"  | "NotIn"  | "Exists"  | "DoesNotExist"  | "Gt"  | "Lt" ;
export const LabelSelectorOperator_In: LabelSelectorOperator = "In";
export const LabelSelectorOperator_NotIn: LabelSelectorOperator = "NotIn";
export const LabelSelectorOperator_Exists: LabelSelectorOperator = "Exists";
export const LabelSelectorOperator_DoesNotExist: LabelSelectorOperator = "DoesNotExist";
export const LabelSelectorOperator_Gt: LabelSelectorOperator = "Gt";
export const LabelSelectorOperator_Lt: LabelSelectorOperator = "Lt";

export interface GetProfileRequest {
  id?: string;
  email?: string;
}

export interface GetChildRequest {
  childId?: string;
}

export interface GetChildrenFromProfileRequest {
  profileId?: string;
}

export interface ChangeChildActivationRequest {
  childId?: string;
  active?: boolean;
}

export interface GetChildrenFromProfileResponse {
  children?: apipb_models.Child[];
}

export interface GetGameReleaseRequest {
  gameId?: string;
  version?: string;
  state?: RequestReleaseState;
}

export interface SoundEnableRequest {
  childId?: string;
  profileId?: string;
  enable?: boolean;
}

export interface GameEntryRequest {
  childId?: string;
  gameId?: string;
  type?: GameEntryRequestRequestType;
  settings?: string;
  index?: number;
}

export interface PublishResponse {
  type?: number;
  message?: string;
  token?: apipb_models.UploadToken;
}

export interface ValidateRequest {
  gameId?: string;
  gameVersion?: string;
  newState?: apipb_models.ReleaseState;
}

export interface UpdateIndecesRequest {
  profileId?: string;
  childId?: string;
  gameIds?: string[];
}

/**
Get game by game_id or unique_name
*/
export interface GetGameRequest {
  uniqueName?: string;
  gameId?: string;
}

export interface ListGamesRequest {
  releaseState?: ListGamesRequestInnerState;
  limit?: number;
  offset?: number;
  language?: string;
}

export interface ListItem {
  gameId?: string;
  uniqueName?: string;
  latestVersion?: string;
  latestState?: apipb_models.ReleaseState;
  productionVersion?: string;
  storage?: string;
  archiveFormat?: string;
  releasedAt?: string|number;
  languages?: string[];
}

export interface ListItemWithTests {
  gameId?: string;
  uniqueName?: string;
  testingVersion?: number;
  groups?: ListItemWithTestsTestGroup[];
}

export interface ListItemWithTestsTestGroup {
  name?: string;
  weight?: number;
  latestVersion?: string;
  latestState?: apipb_models.ReleaseState;
  productionVersion?: string;
  storage?: string;
  archiveFormat?: string;
  releasedAt?: string|number;
  languages?: string[];
}

export interface GetLatestVersionsRequest {
  state?: RequestReleaseState;
  gameIds?: string[];
/**
Device Capabilities
*/
  capabilities?: string[];
}

export interface GameAndVersion {
  gameId?: string;
  version?: string;
  tarballUrl?: string;
}

export interface GameVersionsResponse {
  results?: GameAndVersion[];
}

/**
Search Service
*/
export interface IndexRequest {
}

export interface SearchResult {
  gameId?: string;
  score?: number;
}

export interface SearchRequest {
  query?: string;
  state?: RequestReleaseState;
}

export interface SearchResponse {
  type?: number;
  results?: SearchResult[];
}

/**
Common
*/
export interface Response {
  type?: number;
  message?: string;
}

/**
A label selector requirement is a selector that contains values, a key, and an operator
that relates the key and values.
*/
export interface LabelSelectorRequirement {
/**
key is the label key that the selector applies to.
*/
  key?: string;
/**
operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
*/
  operator?: LabelSelectorOperator;
/**
values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. If the operator is Gt or Lt, the values
array must have a single element, which will be interpreted as an integer.
*/
  values?: string[];
}

/**
An empty label selector term matches all objects. A null label selector term
matches no objects.
*/
export interface LabelSelectorTerm {
/**
expressions is a list of label selector requirements. The requirements are ANDed.
*/
  expressions?: LabelSelectorRequirement[];
}

/**
A label selector represents the union of the results of one or more label queries
over a set of labels; that is, it represents the OR of the selectors represented
by the labelSelectorTerms.
*/
export interface LabelSelector {
/**
terms is a list of label selector terms. The terms are ORed.
*/
  terms?: LabelSelectorTerm[];
}
