/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: ListVideos
// ====================================================

export interface ListVideos_videos {
  __typename: "Video";
  id: string;
  cameraName: string;
  date: any;
  duration: number;
}

export interface ListVideos {
  videos: (ListVideos_videos | null)[] | null;
}

export interface ListVideosVariables {
  date?: any | null;
}
