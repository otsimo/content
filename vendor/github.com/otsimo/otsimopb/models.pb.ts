// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!


export type Gender =  "UNKNOWN"  | "MALE"  | "FEMALE" ;
export const Gender_UNKNOWN: Gender = "UNKNOWN";
export const Gender_MALE: Gender = "MALE";
export const Gender_FEMALE: Gender = "FEMALE";

export type ReleaseState =  "CREATED"  | "DEVELOPMENT"  | "WAITING"  | "REJECTED"  | "VALIDATED"  | "PRODUCTION" ;
export const ReleaseState_CREATED: ReleaseState = "CREATED";
export const ReleaseState_DEVELOPMENT: ReleaseState = "DEVELOPMENT";
export const ReleaseState_WAITING: ReleaseState = "WAITING";
export const ReleaseState_REJECTED: ReleaseState = "REJECTED";
export const ReleaseState_VALIDATED: ReleaseState = "VALIDATED";
export const ReleaseState_PRODUCTION: ReleaseState = "PRODUCTION";

export interface Address {
  streetAddress?: string;
  city?: string;
  state?: string;
  zipCode?: string;
  countryCode?: string;
}

export interface Profile {
  id?: string;
  email?: string;
  firstName?: string;
  lastName?: string;
  language?: string;
  mobilePhone?: string;
  address?: Address;
  createdAt?: string|number;
  lastUpdated?: string|number;
  locale?: string;
  timezone?: number;
  country?: string;
  labels?: { [key: string]: string };
		}

export interface ChildGameEntry {
  id?: string;
  active?: boolean;
  dashboardIndex?: number;
  settings?: string;
  addedAt?: string|number;
  activationChangedAt?: string|number;
  updatedAt?: string|number;
  localSettings?: string;
  localSettingsVersion?: string|number;
}

export interface Child {
  id?: string;
  parentId?: string;
  firstName?: string;
  lastName?: string;
  birthDay?: string|number;
  gender?: Gender;
  language?: string;
  games?: ChildGameEntry[];
  active?: boolean;
  loggedIn?: boolean;
  soundsEnabled?: boolean;
  locale?: string;
  createdAt?: string|number;
  lastUpdated?: string|number;
  childInfo?: { [key: string]: string };
		  totalStarCount?: number;
  levelStarCount?: number;
  level?: number;
  badges?: Badge[];
}

export interface Badge {
  name?: string;
  createdAt?: string|number;
  level?: number;
}

export interface Author {
  name?: string;
  email?: string;
}

export interface GameMetadata {
  language?: string;
/**
Visible Name
*/
  visibleName?: string;
/**
Summary is summary of game
*/
  summary?: string;
/**
Description
*/
  description?: string;
/**
Logo is a rectangle image
*/
  logo?: string;
/**
Icon is a square image
*/
  icon?: string;
/**
Keywords
*/
  keywords?: string[];
/**
Images are image files that show on market
*/
  images?: string[];
/**
InfoSlug is the id of the content which describes how to play this game
*/
  infoSlug?: string;
/**
Assets are asset folders or paths for this language
*/
  assets?: string[];
/**
Localized Annotations
*/
  annotations?: { [key: string]: string };
		}

export interface GameManifest {
/**
Unique Name of game
*/
  uniqueName?: string;
/**
Licence
*/
  licence?: string;
/**
Languages
*/
  languages?: string[];
/**
Homepage is the website of game
*/
  homepage?: string;
/**
Main points to index.html file
*/
  main?: string;
/**
Version is current version for using on manifest file
*/
  version?: string;
/**
Authors is for using at manifest file
*/
  authors?: Author[];
/**
Repository
*/
  repository?: string;
/**
SupportedOrientations are the orientations that can be used for this app
*/
  supportedOrientations?: string[];
/**
Metadata information for each languages
*/
  metadata?: GameMetadata[];
/**
Exclude directories when building
*/
  exclude?: string[];
/**
Settings.json path
*/
  settings?: string;
/**
KV directory path
*/
  kvPath?: string;
/**
DeveloperName is the visible developer name
*/
  developerName?: string;
/**
DefaultLanguage
*/
  defaultLanguage?: string;
/**
Loading Background Color of the Game
*/
  loadingColor?: string;
/**
Capabilities are required host app features. The most basic capability is 'sandbox'.
*/
  capabilities?: string[];
/**
AbTest determines whether or not this version is for ab testing
*/
  abTest?: string;
/**
Labels are string key value pairs
*/
  labels?: { [key: string]: string };
		/**
ManifestVersion
*/
  manifestVersion?: number;
  options?: { [key: string]: GameOption };
		}

export interface GameOption {
  id?: string;
  type?: string;
  description?: string;
  default?: string;
  enum?: string[];
  format?: string;
  minValue?: number;
  maxValue?: number;
}

export interface Game {
  id?: string;
  uniqueName?: string;
  ownerId?: string;
  productionVersion?: string;
  isOnProduction?: boolean;
  createdAt?: string|number;
  lastUpdated?: string|number;
  labels?: { [key: string]: string };
		}

export interface GameRelease {
  releaseId?: string;
  gameId?: string;
  version?: string;
  gameManifest?: GameManifest;
  releasedAt?: string|number;
  releasedBy?: string;
  releaseState?: ReleaseState;
  validatedBy?: string;
  validatedAt?: string|number;
  intVersion?: string|number;
  storage?: string;
  archiveFormat?: string;
  packageUrl?: string;
}

export interface UploadToken {
  token?: string;
  expiresAt?: string|number;
  userId?: string;
  gameId?: string;
  newVersion?: string;
  uploadTo?: string;
}

