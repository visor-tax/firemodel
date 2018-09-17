// DO NOT EDIT - Code generated by firemodel dev.

import { firestore } from 'firebase';

export interface DocumentSnapshot<DataType = firestore.DocumentData>
  extends firestore.DocumentSnapshot {
  data(options?: firestore.SnapshotOptions): DataType | undefined;
}
export interface QueryDocumentSnapshot<
  DataType = firestore.DocumentData
> extends firestore.QueryDocumentSnapshot {
  data(options?: firestore.SnapshotOptions): DataType | undefined;
}
export interface QuerySnapshot<DataType = firestore.DocumentData>
  extends firestore.QuerySnapshot {
  readonly docs: QueryDocumentSnapshot<DataType>[];
}
export interface DocumentSnapshotExpanded<
  DataType = firestore.DocumentData
> {
  exists: firestore.DocumentSnapshot['exists'];
  ref: firestore.DocumentSnapshot['ref'];
  id: firestore.DocumentSnapshot['id'];
  metadata: firestore.DocumentSnapshot['metadata'];
  data: DataType;
}
export interface QuerySnapshotExpanded<
  DataType = firestore.DocumentData
> {
  metadata: {
    hasPendingWrites: firestore.QuerySnapshot['metadata']['hasPendingWrites'];
    fromCache: firestore.QuerySnapshot['metadata']['fromCache'];
  };
  size: firestore.QuerySnapshot['size'];
  empty: firestore.QuerySnapshot['empty'];
  docs: {
    [docId: string]: DocumentSnapshotExpanded<DataType>;
  };
}
export interface DocumentReference<DataType> extends firestore.DocumentReference {
  data(options?: firestore.SnapshotOptions): DataType | undefined;
  get(options?: firestore.GetOptions): Promise<DocumentSnapshot<DataType>>;
}
export interface CollectionReference<
  DataType = firestore.DocumentData
> extends firestore.CollectionReference {
  get(options?: firestore.GetOptions): Promise<QuerySnapshot<DataType>>;
}
export interface Collection<DataType = firestore.DocumentData> {
  [id: string]: DocumentSnapshotExpanded<DataType>;
}

// tslint:disable-next-line:no-namespace
export namespace example {
  type URL = string;

  export interface IFile {
    url: URL;
    mimeType: string;
    name: string;
	}

  /** TODO: Add documentation to TestDirection. */
  export enum TestDirection {
    /** TODO: Add documentation to left. */
    left = "LEFT",
    /** TODO: Add documentation to right. */
    right = "RIGHT",
    /** TODO: Add documentation to up. */
    up = "UP",
    /** TODO: Add documentation to down. */
    down = "DOWN",
  }

  /** A Test is a test model. */
  export interface ITestModel {
    /** TODO: Add documentation to nested_collection. */
    nestedCollection: CollectionReference<ITestModel>;
    /** The name. */
    name?: string;
    /** The age. */
    age?: number;
    /** The number pi. */
    pi?: number;
    /** The birth date. */
    birthdate?: firestore.Timestamp;
    /** True if it is good. */
    isGood?: boolean;
    /** TODO: Add documentation to data. */
    data?: firestore.Blob;
    /** TODO: Add documentation to friend. */
    friend?: DocumentReference<ITestModel>;
    /** TODO: Add documentation to location. */
    location?: firestore.GeoPoint;
    /** TODO: Add documentation to colors. */
    colors?: string[];
    /** TODO: Add documentation to directions. */
    directions?: TestDirection[];
    /** TODO: Add documentation to models. */
    models?: ITestModel[];
    /** TODO: Add documentation to refs. */
    refs?: firestore.DocumentReference[];
    /** TODO: Add documentation to meta. */
    meta?: { [key: string]: any; };
    /** TODO: Add documentation to meta_strs. */
    metaStrs?: { [key: string]: string; };
    /** TODO: Add documentation to direction. */
    direction?: TestDirection;
    /** TODO: Add documentation to test_file. */
    testFile?: IFile;
    /** TODO: Add documentation to url. */
    url?: URL;
    /** TODO: Add documentation to nested. */
    nested?: ITestModel;

    /** Record creation timestamp. */
    createdAt?: firestore.Timestamp;
    /** Record update timestamp. */
    updatedAt?: firestore.Timestamp;
  }

  /** TODO: Add documentation to TestTimestamps. */
  export interface ITestTimestamps {

    /** Record creation timestamp. */
    createdAt?: firestore.Timestamp;
    /** Record update timestamp. */
    updatedAt?: firestore.Timestamp;
  }
}
