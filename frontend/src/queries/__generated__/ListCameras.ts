/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: ListCameras
// ====================================================

export interface ListCameras_cameras {
  __typename: "Camera";
  id: string;
  name: string;
}

export interface ListCameras {
  cameras: (ListCameras_cameras | null)[] | null;
}
