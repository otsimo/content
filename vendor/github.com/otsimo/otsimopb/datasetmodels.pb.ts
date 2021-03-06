// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!


export type ColumnType =  "STRING"  | "INTEGER"  | "REAL"  | "DATE"  | "DATE_TIME"  | "TIME_OF_DAY" ;
export const ColumnType_STRING: ColumnType = "STRING";
export const ColumnType_INTEGER: ColumnType = "INTEGER";
export const ColumnType_REAL: ColumnType = "REAL";
export const ColumnType_DATE: ColumnType = "DATE";
export const ColumnType_DATE_TIME: ColumnType = "DATE_TIME";
export const ColumnType_TIME_OF_DAY: ColumnType = "TIME_OF_DAY";

export interface Column {
  type?: ColumnType;
  name?: string;
}

export interface TimeOfDay {
  hours?: number;
  minutes?: number;
  seconds?: number;
  milliseconds?: number;
}

export interface RowValue {
  str?: string;
  int?: number;
  real?: number;
  date?: string|number;
  dateOfTime?: string|number;
  timeOfDay?: TimeOfDay;
}

export interface Row {
  values?: RowValue[];
}

export interface DataSet {
  label?: string;
  columns?: Column[];
  rows?: Row[];
}

