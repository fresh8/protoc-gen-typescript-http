// Code generated by protoc-gen-typescript-http. DO NOT EDIT.
/* eslint-disable camelcase */

// A shipment represents transportation of goods between an origin
// [site][einride.example.freight.v1.Site] and a destination
// [site][einride.example.freight.v1.Site].
export type Shipment = {
  // The resource name of the shipment.
  name: string | undefined;
  // The creation timestamp of the shipment.
  //
  // Behaviors: OUTPUT_ONLY
  createTime: wellKnownTimestamp | undefined;
  // The last update timestamp of the shipment.
  // Updated when create/update/delete operation is shipment.
  //
  // Behaviors: OUTPUT_ONLY
  updateTime: wellKnownTimestamp | undefined;
  // The deletion timestamp of the shipment.
  //
  // Behaviors: OUTPUT_ONLY
  deleteTime: wellKnownTimestamp | undefined;
  // The resource name of the origin site of the shipment.
  // Format: shippers/{shipper}/sites/{site}
  //
  // Behaviors: REQUIRED
  originSite: string | undefined;
  // The resource name of the destination site of the shipment.
  // Format: shippers/{shipper}/sites/{site}
  //
  // Behaviors: REQUIRED
  destinationSite: string | undefined;
  // The earliest pickup time of the shipment at the origin site.
  //
  // Behaviors: REQUIRED
  pickupEarliestTime: wellKnownTimestamp | undefined;
  // The latest pickup time of the shipment at the origin site.
  //
  // Behaviors: REQUIRED
  pickupLatestTime: wellKnownTimestamp | undefined;
  // The earliest delivery time of the shipment at the destination site.
  //
  // Behaviors: REQUIRED
  deliveryEarliestTime: wellKnownTimestamp | undefined;
  // The latest delivery time of the shipment at the destination site.
  //
  // Behaviors: REQUIRED
  deliveryLatestTime: wellKnownTimestamp | undefined;
  // The line items of the shipment.
  lineItems: LineItem[] | undefined;
  // Annotations of the shipment.
  annotations: { [key: string]: string } | undefined;
};

// Encoded using RFC 3339, where generated output will always be Z-normalized
// and uses 0, 3, 6 or 9 fractional digits.
// Offsets other than "Z" are also accepted.
type wellKnownTimestamp = string;

// A shipment line item.
export type LineItem = {
  // The title of the line item.
  title: string | undefined;
  // The quantity of the line item.
  quantity: number | undefined;
  // The weight of the line item in kilograms.
  weightKg: number | undefined;
  // The volume of the line item in cubic meters.
  volumeM3: number | undefined;
};

// A shipper is a supplier or owner of goods to be transported.
export type Shipper = {
  // The resource name of the shipper.
  name: string | undefined;
  // The creation timestamp of the shipper.
  //
  // Behaviors: OUTPUT_ONLY
  createTime: wellKnownTimestamp | undefined;
  // The last update timestamp of the shipper.
  // Updated when create/update/delete operation is performed.
  //
  // Behaviors: OUTPUT_ONLY
  updateTime: wellKnownTimestamp | undefined;
  // The deletion timestamp of the shipper.
  //
  // Behaviors: OUTPUT_ONLY
  deleteTime: wellKnownTimestamp | undefined;
  // The display name of the shipper.
  //
  // Behaviors: REQUIRED
  displayName: string | undefined;
};

// A site is a node in a [shipper][einride.example.freight.v1.Shipper]'s
// transport network.
export type Site = {
  // The resource name of the site.
  name: string | undefined;
  // The creation timestamp of the site.
  //
  // Behaviors: OUTPUT_ONLY
  createTime: wellKnownTimestamp | undefined;
  // The last update timestamp of the site.
  // Updated when create/update/delete operation is performed.
  //
  // Behaviors: OUTPUT_ONLY
  updateTime: wellKnownTimestamp | undefined;
  // The deletion timestamp of the site.
  //
  // Behaviors: OUTPUT_ONLY
  deleteTime: wellKnownTimestamp | undefined;
  // The display name of the site.
  //
  // Behaviors: REQUIRED
  displayName: string | undefined;
  // The geographic location of the site.
  latLng: googletype_LatLng | undefined;
};

// An object that represents a latitude/longitude pair. This is expressed as a
// pair of doubles to represent degrees latitude and degrees longitude. Unless
// specified otherwise, this must conform to the
// <a href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// standard</a>. Values must be within normalized ranges.
export type googletype_LatLng = {
  // The latitude in degrees. It must be in the range [-90.0, +90.0].
  latitude: number | undefined;
  // The longitude in degrees. It must be in the range [-180.0, +180.0].
  longitude: number | undefined;
};

// Request message for FreightService.GetShipper.
export type GetShipperRequest = {
  // The resource name of the shipper to retrieve.
  // Format: shippers/{shipper}
  //
  // Behaviors: REQUIRED
  name: string | undefined;
};

// Request message for FreightService.ListShippers.
export type ListShippersRequest = {
  // Requested page size. Server may return fewer shippers than requested.
  // If unspecified, server will pick an appropriate default.
  pageSize: number | undefined;
  // A token identifying a page of results the server should return.
  // Typically, this is the value of
  // [ListShippersResponse.next_page_token][einride.example.freight.v1.ListShippersResponse.next_page_token]
  // returned from the previous call to `ListShippers` method.
  pageToken: string | undefined;
};

// Response message for FreightService.ListShippers.
export type ListShippersResponse = {
  // The list of shippers.
  shippers: Shipper[] | undefined;
  // A token to retrieve next page of results.  Pass this value in the
  // [ListShippersRequest.page_token][einride.example.freight.v1.ListShippersRequest.page_token]
  // field in the subsequent call to `ListShippers` method to retrieve the next
  // page of results.
  nextPageToken: string | undefined;
};

// Request message for FreightService.CreateShipper.
export type CreateShipperRequest = {
  // The shipper to create.
  //
  // Behaviors: REQUIRED
  shipper: Shipper | undefined;
};

// Request message for FreightService.UpdateShipper.
export type UpdateShipperRequest = {
  // The shipper to update with. The name must match or be empty.
  // The shipper's `name` field is used to identify the shipper to be updated.
  // Format: shippers/{shipper}
  //
  // Behaviors: REQUIRED
  shipper: Shipper | undefined;
  // The list of fields to be updated.
  updateMask: wellKnownFieldMask | undefined;
};

// In JSON, a field mask is encoded as a single string where paths are
// separated by a comma. Fields name in each path are converted
// to/from lower-camel naming conventions.
// As an example, consider the following message declarations:
//
//     message Profile {
//       User user = 1;
//       Photo photo = 2;
//     }
//     message User {
//       string display_name = 1;
//       string address = 2;
//     }
//
// In proto a field mask for `Profile` may look as such:
//
//     mask {
//       paths: "user.display_name"
//       paths: "photo"
//     }
//
// In JSON, the same mask is represented as below:
//
//     {
//       mask: "user.displayName,photo"
//     }
type wellKnownFieldMask = string;

// Request message for FreightService.DeleteShipper.
export type DeleteShipperRequest = {
  // The resource name of the shipper to delete.
  // Format: shippers/{shipper}
  //
  // Behaviors: REQUIRED
  name: string | undefined;
};

// Request message for FreightService.GetSite.
export type GetSiteRequest = {
  // The resource name of the site to retrieve.
  // Format: shippers/{shipper}/sites/{site}
  //
  // Behaviors: REQUIRED
  name: string | undefined;
};

// Request message for FreightService.ListSites.
export type ListSitesRequest = {
  // The resource name of the parent, which owns this collection of sites.
  // Format: shippers/{shipper}
  //
  // Behaviors: REQUIRED
  parent: string | undefined;
  // Requested page size. Server may return fewer sites than requested.
  // If unspecified, server will pick an appropriate default.
  pageSize: number | undefined;
  // A token identifying a page of results the server should return.
  // Typically, this is the value of
  // [ListSitesResponse.next_page_token][einride.example.freight.v1.ListSitesResponse.next_page_token]
  // returned from the previous call to `ListSites` method.
  pageToken: string | undefined;
};

// Response message for FreightService.ListSites.
export type ListSitesResponse = {
  // The list of sites.
  sites: Site[] | undefined;
  // A token to retrieve next page of results.  Pass this value in the
  // [ListSitesRequest.page_token][einride.example.freight.v1.ListSitesRequest.page_token]
  // field in the subsequent call to `ListSites` method to retrieve the next
  // page of results.
  nextPageToken: string | undefined;
};

// Request message for FreightService.CreateSite.
export type CreateSiteRequest = {
  // The resource name of the parent shipper for which this site will be created.
  // Format: shippers/{shipper}
  //
  // Behaviors: REQUIRED
  parent: string | undefined;
  // The site to create.
  //
  // Behaviors: REQUIRED
  site: Site | undefined;
};

// Request message for FreightService.UpdateSite.
export type UpdateSiteRequest = {
  // The site to update with. The name must match or be empty.
  // The site's `name` field is used to identify the site to be updated.
  // Format: shippers/{shipper}/sites/{site}
  //
  // Behaviors: REQUIRED
  site: Site | undefined;
  // The list of fields to be updated.
  updateMask: wellKnownFieldMask | undefined;
};

// Request message for FreightService.DeleteSite.
export type DeleteSiteRequest = {
  // The resource name of the site to delete.
  // Format: shippers/{shipper}/sites/{site}
  //
  // Behaviors: REQUIRED
  name: string | undefined;
};

// Request message for FreightService.GetShipment.
export type GetShipmentRequest = {
  // The resource name of the shipment to retrieve.
  // Format: shippers/{shipper}/shipments/{shipment}
  //
  // Behaviors: REQUIRED
  name: string | undefined;
};

// Request message for FreightService.ListShipments.
export type ListShipmentsRequest = {
  // The resource name of the parent, which owns this collection of shipments.
  // Format: shippers/{shipper}
  //
  // Behaviors: REQUIRED
  parent: string | undefined;
  // Requested page size. Server may return fewer shipments than requested.
  // If unspecified, server will pick an appropriate default.
  pageSize: number | undefined;
  // A token identifying a page of results the server should return.
  // Typically, this is the value of
  // [ListShipmentsResponse.next_page_token][einride.example.freight.v1.ListShipmentsResponse.next_page_token]
  // returned from the previous call to `ListShipments` method.
  pageToken: string | undefined;
};

// Response message for FreightService.ListShipments.
export type ListShipmentsResponse = {
  // The list of shipments.
  shipments: Shipment[] | undefined;
  // A token to retrieve next page of results.  Pass this value in the
  // [ListShipmentsRequest.page_token][einride.example.freight.v1.ListShipmentsRequest.page_token]
  // field in the subsequent call to `ListShipments` method to retrieve the next
  // page of results.
  nextPageToken: string | undefined;
};

// Request message for FreightService.CreateShipment.
export type CreateShipmentRequest = {
  // The resource name of the parent shipper for which this shipment will be created.
  // Format: shippers/{shipper}
  //
  // Behaviors: REQUIRED
  parent: string | undefined;
  // The shipment to create.
  //
  // Behaviors: REQUIRED
  shipment: Shipment | undefined;
};

// Request message for FreightService.UpdateShipment.
export type UpdateShipmentRequest = {
  // The shipment to update with. The name must match or be empty.
  // The shipment's `name` field is used to identify the shipment to be updated.
  // Format: shippers/{shipper}/shipments/{shipment}
  //
  // Behaviors: REQUIRED
  shipment: Shipment | undefined;
  // The list of fields to be updated.
  updateMask: wellKnownFieldMask | undefined;
};

// Request message for FreightService.DeleteShipment.
export type DeleteShipmentRequest = {
  // The resource name of the shipment to delete.
  // Format: shippers/{shipper}/shipments/{shipment}
  //
  // Behaviors: REQUIRED
  name: string | undefined;
};

// This API represents a simple freight service.
// It defines the following resource model:
// - The API has a collection of [Shipper][einride.example.freight.v1.Shipper]
// resources, named `shippers/*`
// - Each Shipper has a collection of [Site][einride.example.freight.v1.Site]
// resources, named `shippers/*/sites/*`
// - Each Shipper has a collection of [Shipment][einride.example.freight.v1.Shipment]
// resources, named `shippers/*/shipments/*`
export interface FreightService {
  // Get a shipper.
  // See: https://google.aip.dev/131 (Standard methods: Get).
  GetShipper(request: GetShipperRequest): Promise<Shipper>;
  // List shippers.
  // See: https://google.aip.dev/132 (Standard methods: List).
  ListShippers(request: ListShippersRequest): Promise<ListShippersResponse>;
  // Create a shipper.
  // See: https://google.aip.dev/133 (Standard methods: Create).
  CreateShipper(request: CreateShipperRequest): Promise<Shipper>;
  // Update a shipper.
  // See: https://google.aip.dev/134 (Standard methods: Update).
  UpdateShipper(request: UpdateShipperRequest): Promise<Shipper>;
  // Delete a shipper.
  // See: https://google.aip.dev/135 (Standard methods: Delete).
  // See: https://google.aip.dev/164 (Soft delete).
  DeleteShipper(request: DeleteShipperRequest): Promise<Shipper>;
  // Get a site.
  // See: https://google.aip.dev/131 (Standard methods: Get).
  GetSite(request: GetSiteRequest): Promise<Site>;
  // List sites for a shipper.
  // See: https://google.aip.dev/132 (Standard methods: List).
  ListSites(request: ListSitesRequest): Promise<ListSitesResponse>;
  // Create a site.
  // See: https://google.aip.dev/133 (Standard methods: Create).
  CreateSite(request: CreateSiteRequest): Promise<Site>;
  // Update a site.
  // See: https://google.aip.dev/134 (Standard methods: Update).
  UpdateSite(request: UpdateSiteRequest): Promise<Site>;
  // Delete a site.
  // See: https://google.aip.dev/135 (Standard methods: Delete).
  // See: https://google.aip.dev/164 (Soft delete).
  DeleteSite(request: DeleteSiteRequest): Promise<Site>;
  // Get a shipment.
  // See: https://google.aip.dev/131 (Standard methods: Get).
  GetShipment(request: GetShipmentRequest): Promise<Shipment>;
  // List shipments for a shipper.
  // See: https://google.aip.dev/132 (Standard methods: List).
  ListShipments(request: ListShipmentsRequest): Promise<ListShipmentsResponse>;
  // Create a shipment.
  // See: https://google.aip.dev/133 (Standard methods: Create).
  CreateShipment(request: CreateShipmentRequest): Promise<Shipment>;
  // Update a shipment.
  // See: https://google.aip.dev/134 (Standard methods: Update).
  UpdateShipment(request: UpdateShipmentRequest): Promise<Shipment>;
  // Delete a shipment.
  // See: https://google.aip.dev/135 (Standard methods: Delete).
  // See: https://google.aip.dev/164 (Soft delete).
  DeleteShipment(request: DeleteShipmentRequest): Promise<Shipment>;
}

type Request = {
  path: string;
  method: string;
  body: string | null;
};

type RequestHandler = (request: Request) => Promise<unknown>;

export function createFreightServiceClient(
  handler: RequestHandler
): FreightService {
  return {
    GetShipper(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.name) {
        throw new Error("missing required field request.name");
      }
      const path = `v1/${request.name}`; // eslint-disable-line quotes
      const body = null;
      const queryParams: string[] = [];
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "GET",
        body,
      }) as Promise<Shipper>;
    },
    ListShippers(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      const path = `v1/shippers`; // eslint-disable-line quotes
      const body = null;
      const queryParams: string[] = [];
      if (request.pageSize) {
        queryParams.push(`pageSize=${encodeURIComponent(request.pageSize.toString())}`)
      }
      if (request.pageToken) {
        queryParams.push(`pageToken=${encodeURIComponent(request.pageToken.toString())}`)
      }
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "GET",
        body,
      }) as Promise<ListShippersResponse>;
    },
    CreateShipper(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      const path = `v1/shippers`; // eslint-disable-line quotes
      const body = JSON.stringify(request?.shipper ?? {});
      const queryParams: string[] = [];
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "POST",
        body,
      }) as Promise<Shipper>;
    },
    UpdateShipper(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.shipper?.name) {
        throw new Error("missing required field request.shipper.name");
      }
      const path = `v1/${request.shipper.name}`; // eslint-disable-line quotes
      const body = JSON.stringify(request?.shipper ?? {});
      const queryParams: string[] = [];
      if (request.updateMask) {
        queryParams.push(`updateMask=${encodeURIComponent(request.updateMask.toString())}`)
      }
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "PATCH",
        body,
      }) as Promise<Shipper>;
    },
    DeleteShipper(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.name) {
        throw new Error("missing required field request.name");
      }
      const path = `v1/${request.name}`; // eslint-disable-line quotes
      const body = null;
      const queryParams: string[] = [];
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "DELETE",
        body,
      }) as Promise<Shipper>;
    },
    GetSite(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.name) {
        throw new Error("missing required field request.name");
      }
      const path = `v1/${request.name}`; // eslint-disable-line quotes
      const body = null;
      const queryParams: string[] = [];
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "GET",
        body,
      }) as Promise<Site>;
    },
    ListSites(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.parent) {
        throw new Error("missing required field request.parent");
      }
      const path = `v1/${request.parent}/sites`; // eslint-disable-line quotes
      const body = null;
      const queryParams: string[] = [];
      if (request.pageSize) {
        queryParams.push(`pageSize=${encodeURIComponent(request.pageSize.toString())}`)
      }
      if (request.pageToken) {
        queryParams.push(`pageToken=${encodeURIComponent(request.pageToken.toString())}`)
      }
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "GET",
        body,
      }) as Promise<ListSitesResponse>;
    },
    CreateSite(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.parent) {
        throw new Error("missing required field request.parent");
      }
      const path = `v1/${request.parent}/sites`; // eslint-disable-line quotes
      const body = JSON.stringify(request?.site ?? {});
      const queryParams: string[] = [];
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "POST",
        body,
      }) as Promise<Site>;
    },
    UpdateSite(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.site?.name) {
        throw new Error("missing required field request.site.name");
      }
      const path = `v1/${request.site.name}`; // eslint-disable-line quotes
      const body = JSON.stringify(request?.site ?? {});
      const queryParams: string[] = [];
      if (request.updateMask) {
        queryParams.push(`updateMask=${encodeURIComponent(request.updateMask.toString())}`)
      }
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "PATCH",
        body,
      }) as Promise<Site>;
    },
    DeleteSite(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.name) {
        throw new Error("missing required field request.name");
      }
      const path = `v1/${request.name}`; // eslint-disable-line quotes
      const body = null;
      const queryParams: string[] = [];
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "DELETE",
        body,
      }) as Promise<Site>;
    },
    GetShipment(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.name) {
        throw new Error("missing required field request.name");
      }
      const path = `v1/${request.name}`; // eslint-disable-line quotes
      const body = null;
      const queryParams: string[] = [];
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "GET",
        body,
      }) as Promise<Shipment>;
    },
    ListShipments(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.parent) {
        throw new Error("missing required field request.parent");
      }
      const path = `v1/${request.parent}/shipments`; // eslint-disable-line quotes
      const body = null;
      const queryParams: string[] = [];
      if (request.pageSize) {
        queryParams.push(`pageSize=${encodeURIComponent(request.pageSize.toString())}`)
      }
      if (request.pageToken) {
        queryParams.push(`pageToken=${encodeURIComponent(request.pageToken.toString())}`)
      }
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "GET",
        body,
      }) as Promise<ListShipmentsResponse>;
    },
    CreateShipment(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.parent) {
        throw new Error("missing required field request.parent");
      }
      const path = `v1/${request.parent}/shipments`; // eslint-disable-line quotes
      const body = JSON.stringify(request?.shipment ?? {});
      const queryParams: string[] = [];
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "POST",
        body,
      }) as Promise<Shipment>;
    },
    UpdateShipment(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.shipment?.name) {
        throw new Error("missing required field request.shipment.name");
      }
      const path = `v1/${request.shipment.name}`; // eslint-disable-line quotes
      const body = JSON.stringify(request?.shipment ?? {});
      const queryParams: string[] = [];
      if (request.updateMask) {
        queryParams.push(`updateMask=${encodeURIComponent(request.updateMask.toString())}`)
      }
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "PATCH",
        body,
      }) as Promise<Shipment>;
    },
    DeleteShipment(request) { // eslint-disable-line @typescript-eslint/no-unused-vars
      if (!request.name) {
        throw new Error("missing required field request.name");
      }
      const path = `v1/${request.name}`; // eslint-disable-line quotes
      const body = null;
      const queryParams: string[] = [];
      let uri = path;
      if (queryParams.length > 0) {
        uri += `?${queryParams.join("&")}`
      }
      return handler({
        path: uri,
        method: "DELETE",
        body,
      }) as Promise<Shipment>;
    },
  };
}

// @@protoc_insertion_point(typescript-http-eof)
