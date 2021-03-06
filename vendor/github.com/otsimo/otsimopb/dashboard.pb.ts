// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!

import * as apipb_datasetmodels from "./datasetmodels.pb";

export type CardDecorationSize =  "SMALL"  | "MEDIUM"  | "LARGE" ;
/**
Small is 1x1 block on iphone
*/
export const CardDecorationSize_SMALL: CardDecorationSize = "SMALL";
/**
Medium is 2x1 block on iphone
*/
export const CardDecorationSize_MEDIUM: CardDecorationSize = "MEDIUM";
/**
Large is 2x2 block on iphone
*/
export const CardDecorationSize_LARGE: CardDecorationSize = "LARGE";

export type CardDecorationBackgroundStyle =  "EMPTY"  | "IMAGE"  | "CHART_SILHOUETTE" ;
export const CardDecorationBackgroundStyle_EMPTY: CardDecorationBackgroundStyle = "EMPTY";
export const CardDecorationBackgroundStyle_IMAGE: CardDecorationBackgroundStyle = "IMAGE";
export const CardDecorationBackgroundStyle_CHART_SILHOUETTE: CardDecorationBackgroundStyle = "CHART_SILHOUETTE";

export type ChartType =  "LINE"  | "BAR"  | "PIE"  | "SCATTER"  | "BUBLE"  | "RADAR"  | "GEO"  | "TIMELINE" ;
export const ChartType_LINE: ChartType = "LINE";
export const ChartType_BAR: ChartType = "BAR";
export const ChartType_PIE: ChartType = "PIE";
export const ChartType_SCATTER: ChartType = "SCATTER";
export const ChartType_BUBLE: ChartType = "BUBLE";
export const ChartType_RADAR: ChartType = "RADAR";
export const ChartType_GEO: ChartType = "GEO";
export const ChartType_TIMELINE: ChartType = "TIMELINE";

export interface DashboardItems {
/**
ProfileId
*/
  profileId?: string;
/**
ChildId
*/
  childId?: string;
/**
Created At
*/
  createdAt?: string|number;
  items?: Card[];
}

export interface DashboardGetRequest {
  profileId?: string;
  childId?: string;
  appVersion?: string;
  language?: string;
  countryCode?: string;
  lastTimeDataFetched?: string|number;
}

export interface CardDecoration {
  size?: CardDecorationSize;
  backgroundStyle?: CardDecorationBackgroundStyle;
  imageUrl?: string;
  leftIcon?: string;
  rightIcon?: string;
}

export interface CardEmpty {
}

export interface CardWebpage {
  url?: string;
}

export interface CardApplink {
  applink?: string;
}

export interface CardAnalysis {
  data?: apipb_datasetmodels.DataSet;
  chartType?: ChartType;
}

export interface Card {
  id?: string;
  text?: string;
  expiresAt?: string|number;
  createdAt?: string|number;
  decoration?: CardDecoration;
/**
Score is between 0-500
*/
  providerScore?: number;
/**
ProviderWeight is between 0-2
*/
  providerWeight?: number;
  providerName?: string;
  language?: string;
  empty?: CardEmpty;
  webpage?: CardWebpage;
  applink?: CardApplink;
  analysis?: CardAnalysis;
/**
Title for newer systems
*/
  title?: string;
/**
Subtitle for newer systems
*/
  subtitle?: string;
/**
Labels of the card
*/
  labels?: { [key: string]: string };
		}

