// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!

import * as otsimo_dashboard from "./dashboard.pb";

export interface ProviderGetRequest {
  request?: otsimo_dashboard.DashboardGetRequest;
  userGroups?: string[];
}

export interface ProviderItem {
  cacheable?: boolean;
  ttl?: string|number;
  item?: otsimo_dashboard.Card;
}

export interface ProviderItems {
/**
ProfileId
*/
  profileId?: string;
/**
ChildId
*/
  childId?: string;
/**
CreatedAt
*/
  createdAt?: string|number;
/**
Cacheable
*/
  cacheable?: boolean;
/**
TTL is titme to live duration
*/
  ttl?: string|number;
/**
Items
*/
  items?: ProviderItem[];
}

export interface ProviderInfoRequest {
}

export interface ProviderInfo {
  servesFor?: string[];
}

