/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter eslint_disable,add_pb_suffix,server_grpc1,ts_nocheck
// @generated from protobuf file "teleport/accesslist/v1/accesslist_service.proto" (package "teleport.accesslist.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
//
// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
import { GetInheritedGrantsResponse } from "./accesslist_service_pb";
import { GetInheritedGrantsRequest } from "./accesslist_service_pb";
import { GetSuggestedAccessListsResponse } from "./accesslist_service_pb";
import { GetSuggestedAccessListsRequest } from "./accesslist_service_pb";
import { AccessRequestPromoteResponse } from "./accesslist_service_pb";
import { AccessRequestPromoteRequest } from "./accesslist_service_pb";
import { DeleteAccessListReviewRequest } from "./accesslist_service_pb";
import { CreateAccessListReviewResponse } from "./accesslist_service_pb";
import { CreateAccessListReviewRequest } from "./accesslist_service_pb";
import { ListAllAccessListReviewsResponse } from "./accesslist_service_pb";
import { ListAllAccessListReviewsRequest } from "./accesslist_service_pb";
import { ListAccessListReviewsResponse } from "./accesslist_service_pb";
import { ListAccessListReviewsRequest } from "./accesslist_service_pb";
import { UpsertAccessListWithMembersResponse } from "./accesslist_service_pb";
import { UpsertAccessListWithMembersRequest } from "./accesslist_service_pb";
import { DeleteAllAccessListMembersRequest } from "./accesslist_service_pb";
import { DeleteAllAccessListMembersForAccessListRequest } from "./accesslist_service_pb";
import { DeleteStaticAccessListMemberResponse } from "./accesslist_service_pb";
import { DeleteStaticAccessListMemberRequest } from "./accesslist_service_pb";
import { DeleteAccessListMemberRequest } from "./accesslist_service_pb";
import { UpdateAccessListMemberRequest } from "./accesslist_service_pb";
import { UpsertStaticAccessListMemberResponse } from "./accesslist_service_pb";
import { UpsertStaticAccessListMemberRequest } from "./accesslist_service_pb";
import { UpsertAccessListMemberRequest } from "./accesslist_service_pb";
import { GetAccessListOwnersResponse } from "./accesslist_service_pb";
import { GetAccessListOwnersRequest } from "./accesslist_service_pb";
import { GetStaticAccessListMemberResponse } from "./accesslist_service_pb";
import { GetStaticAccessListMemberRequest } from "./accesslist_service_pb";
import { Member } from "./accesslist_pb";
import { GetAccessListMemberRequest } from "./accesslist_service_pb";
import { ListAllAccessListMembersResponse } from "./accesslist_service_pb";
import { ListAllAccessListMembersRequest } from "./accesslist_service_pb";
import { ListAccessListMembersResponse } from "./accesslist_service_pb";
import { ListAccessListMembersRequest } from "./accesslist_service_pb";
import { CountAccessListMembersResponse } from "./accesslist_service_pb";
import { CountAccessListMembersRequest } from "./accesslist_service_pb";
import { GetAccessListsToReviewResponse } from "./accesslist_service_pb";
import { GetAccessListsToReviewRequest } from "./accesslist_service_pb";
import { DeleteAllAccessListsRequest } from "./accesslist_service_pb";
import { Empty } from "../../../google/protobuf/empty_pb";
import { DeleteAccessListRequest } from "./accesslist_service_pb";
import { UpdateAccessListRequest } from "./accesslist_service_pb";
import { UpsertAccessListRequest } from "./accesslist_service_pb";
import { AccessList } from "./accesslist_pb";
import { GetAccessListRequest } from "./accesslist_service_pb";
import { ListAccessListsResponse } from "./accesslist_service_pb";
import { ListAccessListsRequest } from "./accesslist_service_pb";
import { GetAccessListsResponse } from "./accesslist_service_pb";
import { GetAccessListsRequest } from "./accesslist_service_pb";
import type * as grpc from "@grpc/grpc-js";
/**
 * AccessListService provides CRUD methods for Access List resources.
 *
 * @generated from protobuf service teleport.accesslist.v1.AccessListService
 */
export interface IAccessListService extends grpc.UntypedServiceImplementation {
    /**
     * GetAccessLists returns a list of all access lists.
     *
     * @generated from protobuf rpc: GetAccessLists(teleport.accesslist.v1.GetAccessListsRequest) returns (teleport.accesslist.v1.GetAccessListsResponse);
     */
    getAccessLists: grpc.handleUnaryCall<GetAccessListsRequest, GetAccessListsResponse>;
    /**
     * ListAccessLists returns a paginated list of all access lists.
     *
     * @generated from protobuf rpc: ListAccessLists(teleport.accesslist.v1.ListAccessListsRequest) returns (teleport.accesslist.v1.ListAccessListsResponse);
     */
    listAccessLists: grpc.handleUnaryCall<ListAccessListsRequest, ListAccessListsResponse>;
    /**
     * GetAccessList returns the specified access list resource.
     *
     * @generated from protobuf rpc: GetAccessList(teleport.accesslist.v1.GetAccessListRequest) returns (teleport.accesslist.v1.AccessList);
     */
    getAccessList: grpc.handleUnaryCall<GetAccessListRequest, AccessList>;
    /**
     * UpsertAccessList creates or updates an access list resource.
     *
     * @generated from protobuf rpc: UpsertAccessList(teleport.accesslist.v1.UpsertAccessListRequest) returns (teleport.accesslist.v1.AccessList);
     */
    upsertAccessList: grpc.handleUnaryCall<UpsertAccessListRequest, AccessList>;
    /**
     * UpdateAccessList updates an access list resource.
     *
     * @generated from protobuf rpc: UpdateAccessList(teleport.accesslist.v1.UpdateAccessListRequest) returns (teleport.accesslist.v1.AccessList);
     */
    updateAccessList: grpc.handleUnaryCall<UpdateAccessListRequest, AccessList>;
    /**
     * DeleteAccessList hard deletes the specified access list resource.
     *
     * @generated from protobuf rpc: DeleteAccessList(teleport.accesslist.v1.DeleteAccessListRequest) returns (google.protobuf.Empty);
     */
    deleteAccessList: grpc.handleUnaryCall<DeleteAccessListRequest, Empty>;
    /**
     * DeleteAllAccessLists hard deletes all access lists.
     *
     * @generated from protobuf rpc: DeleteAllAccessLists(teleport.accesslist.v1.DeleteAllAccessListsRequest) returns (google.protobuf.Empty);
     */
    deleteAllAccessLists: grpc.handleUnaryCall<DeleteAllAccessListsRequest, Empty>;
    /**
     * GetAccessListsToReview will return access lists that need to be reviewed by
     * the current user.
     *
     * @generated from protobuf rpc: GetAccessListsToReview(teleport.accesslist.v1.GetAccessListsToReviewRequest) returns (teleport.accesslist.v1.GetAccessListsToReviewResponse);
     */
    getAccessListsToReview: grpc.handleUnaryCall<GetAccessListsToReviewRequest, GetAccessListsToReviewResponse>;
    /**
     * CountAccessListMembers returns the count of access list members in an
     * access list.
     *
     * @generated from protobuf rpc: CountAccessListMembers(teleport.accesslist.v1.CountAccessListMembersRequest) returns (teleport.accesslist.v1.CountAccessListMembersResponse);
     */
    countAccessListMembers: grpc.handleUnaryCall<CountAccessListMembersRequest, CountAccessListMembersResponse>;
    /**
     * ListAccessListMembers returns a paginated list of all access list members.
     *
     * @generated from protobuf rpc: ListAccessListMembers(teleport.accesslist.v1.ListAccessListMembersRequest) returns (teleport.accesslist.v1.ListAccessListMembersResponse);
     */
    listAccessListMembers: grpc.handleUnaryCall<ListAccessListMembersRequest, ListAccessListMembersResponse>;
    /**
     * ListAllAccessListMembers returns a paginated list of all access list
     * members for all access lists.
     *
     * @generated from protobuf rpc: ListAllAccessListMembers(teleport.accesslist.v1.ListAllAccessListMembersRequest) returns (teleport.accesslist.v1.ListAllAccessListMembersResponse);
     */
    listAllAccessListMembers: grpc.handleUnaryCall<ListAllAccessListMembersRequest, ListAllAccessListMembersResponse>;
    /**
     * GetAccessListMember returns the specified access list member resource.
     *
     * @generated from protobuf rpc: GetAccessListMember(teleport.accesslist.v1.GetAccessListMemberRequest) returns (teleport.accesslist.v1.Member);
     */
    getAccessListMember: grpc.handleUnaryCall<GetAccessListMemberRequest, Member>;
    /**
     * GetStaticAccessListMember returns the specified access_list_member resource. If returns error
     * if the target access_list is not of type static.  This API is there for the IaC tools to
     * prevent them from making changes to members of dynamic access lists.
     *
     * @generated from protobuf rpc: GetStaticAccessListMember(teleport.accesslist.v1.GetStaticAccessListMemberRequest) returns (teleport.accesslist.v1.GetStaticAccessListMemberResponse);
     */
    getStaticAccessListMember: grpc.handleUnaryCall<GetStaticAccessListMemberRequest, GetStaticAccessListMemberResponse>;
    /**
     * GetAccessListOwners returns a list of all owners in an Access List,
     * including those inherited from nested Access Lists.
     *
     * @generated from protobuf rpc: GetAccessListOwners(teleport.accesslist.v1.GetAccessListOwnersRequest) returns (teleport.accesslist.v1.GetAccessListOwnersResponse);
     */
    getAccessListOwners: grpc.handleUnaryCall<GetAccessListOwnersRequest, GetAccessListOwnersResponse>;
    /**
     * UpsertAccessListMember creates or updates an access list member resource.
     *
     * @generated from protobuf rpc: UpsertAccessListMember(teleport.accesslist.v1.UpsertAccessListMemberRequest) returns (teleport.accesslist.v1.Member);
     */
    upsertAccessListMember: grpc.handleUnaryCall<UpsertAccessListMemberRequest, Member>;
    /**
     * UpsertStaticAccessListMember creates or updates an access_list_member resource. It returns
     * error and does nothing if the target access_list is not of type static. This API is there for
     * the IaC tools to prevent them from making changes to members of dynamic access lists.
     *
     * @generated from protobuf rpc: UpsertStaticAccessListMember(teleport.accesslist.v1.UpsertStaticAccessListMemberRequest) returns (teleport.accesslist.v1.UpsertStaticAccessListMemberResponse);
     */
    upsertStaticAccessListMember: grpc.handleUnaryCall<UpsertStaticAccessListMemberRequest, UpsertStaticAccessListMemberResponse>;
    /**
     * UpdateAccessListMember conditionally updates an access list member resource.
     *
     * @generated from protobuf rpc: UpdateAccessListMember(teleport.accesslist.v1.UpdateAccessListMemberRequest) returns (teleport.accesslist.v1.Member);
     */
    updateAccessListMember: grpc.handleUnaryCall<UpdateAccessListMemberRequest, Member>;
    /**
     * DeleteAccessListMember hard deletes the specified access list member
     * resource.
     *
     * @generated from protobuf rpc: DeleteAccessListMember(teleport.accesslist.v1.DeleteAccessListMemberRequest) returns (google.protobuf.Empty);
     */
    deleteAccessListMember: grpc.handleUnaryCall<DeleteAccessListMemberRequest, Empty>;
    /**
     * DeleteStaticAccessListMember hard deletes the specified access_list_member. It returns error
     * and does nothing if the target access_list is not of static type. This API is there for the
     * IaC tools to prevent them from making changes to members of dynamic access lists.
     *
     * @generated from protobuf rpc: DeleteStaticAccessListMember(teleport.accesslist.v1.DeleteStaticAccessListMemberRequest) returns (teleport.accesslist.v1.DeleteStaticAccessListMemberResponse);
     */
    deleteStaticAccessListMember: grpc.handleUnaryCall<DeleteStaticAccessListMemberRequest, DeleteStaticAccessListMemberResponse>;
    /**
     * DeleteAllAccessListMembers hard deletes all access list members for an
     * access list.
     *
     * @generated from protobuf rpc: DeleteAllAccessListMembersForAccessList(teleport.accesslist.v1.DeleteAllAccessListMembersForAccessListRequest) returns (google.protobuf.Empty);
     */
    deleteAllAccessListMembersForAccessList: grpc.handleUnaryCall<DeleteAllAccessListMembersForAccessListRequest, Empty>;
    /**
     * DeleteAllAccessListMembers hard deletes all access list members for an
     * access list.
     *
     * @generated from protobuf rpc: DeleteAllAccessListMembers(teleport.accesslist.v1.DeleteAllAccessListMembersRequest) returns (google.protobuf.Empty);
     */
    deleteAllAccessListMembers: grpc.handleUnaryCall<DeleteAllAccessListMembersRequest, Empty>;
    /**
     * UpsertAccessListWithMembers creates or updates an access list with members.
     *
     * @generated from protobuf rpc: UpsertAccessListWithMembers(teleport.accesslist.v1.UpsertAccessListWithMembersRequest) returns (teleport.accesslist.v1.UpsertAccessListWithMembersResponse);
     */
    upsertAccessListWithMembers: grpc.handleUnaryCall<UpsertAccessListWithMembersRequest, UpsertAccessListWithMembersResponse>;
    /**
     * ListAccessListReviews will list access list reviews for a particular access
     * list.
     *
     * @generated from protobuf rpc: ListAccessListReviews(teleport.accesslist.v1.ListAccessListReviewsRequest) returns (teleport.accesslist.v1.ListAccessListReviewsResponse);
     */
    listAccessListReviews: grpc.handleUnaryCall<ListAccessListReviewsRequest, ListAccessListReviewsResponse>;
    /**
     * ListAllAccessListReviews will list access list reviews for all access
     * lists.
     *
     * @generated from protobuf rpc: ListAllAccessListReviews(teleport.accesslist.v1.ListAllAccessListReviewsRequest) returns (teleport.accesslist.v1.ListAllAccessListReviewsResponse);
     */
    listAllAccessListReviews: grpc.handleUnaryCall<ListAllAccessListReviewsRequest, ListAllAccessListReviewsResponse>;
    /**
     * CreateAccessListReview will create a new review for an access list. It will
     * also modify the original access list and its members depending on the
     * details of the review.
     *
     * @generated from protobuf rpc: CreateAccessListReview(teleport.accesslist.v1.CreateAccessListReviewRequest) returns (teleport.accesslist.v1.CreateAccessListReviewResponse);
     */
    createAccessListReview: grpc.handleUnaryCall<CreateAccessListReviewRequest, CreateAccessListReviewResponse>;
    /**
     * DeleteAccessListReview will delete an access list review from the backend.
     *
     * @generated from protobuf rpc: DeleteAccessListReview(teleport.accesslist.v1.DeleteAccessListReviewRequest) returns (google.protobuf.Empty);
     */
    deleteAccessListReview: grpc.handleUnaryCall<DeleteAccessListReviewRequest, Empty>;
    /**
     * AccessRequestPromote promotes an access request to an access list.
     *
     * @generated from protobuf rpc: AccessRequestPromote(teleport.accesslist.v1.AccessRequestPromoteRequest) returns (teleport.accesslist.v1.AccessRequestPromoteResponse);
     */
    accessRequestPromote: grpc.handleUnaryCall<AccessRequestPromoteRequest, AccessRequestPromoteResponse>;
    /**
     * GetSuggestedAccessLists returns suggested access lists for an access
     * request.
     *
     * @generated from protobuf rpc: GetSuggestedAccessLists(teleport.accesslist.v1.GetSuggestedAccessListsRequest) returns (teleport.accesslist.v1.GetSuggestedAccessListsResponse);
     */
    getSuggestedAccessLists: grpc.handleUnaryCall<GetSuggestedAccessListsRequest, GetSuggestedAccessListsResponse>;
    /**
     * GetInheritedGrants returns the inherited grants for an access list.
     *
     * @generated from protobuf rpc: GetInheritedGrants(teleport.accesslist.v1.GetInheritedGrantsRequest) returns (teleport.accesslist.v1.GetInheritedGrantsResponse);
     */
    getInheritedGrants: grpc.handleUnaryCall<GetInheritedGrantsRequest, GetInheritedGrantsResponse>;
}
/**
 * @grpc/grpc-js definition for the protobuf service teleport.accesslist.v1.AccessListService.
 *
 * Usage: Implement the interface IAccessListService and add to a grpc server.
 *
 * ```typescript
 * const server = new grpc.Server();
 * const service: IAccessListService = ...
 * server.addService(accessListServiceDefinition, service);
 * ```
 */
export const accessListServiceDefinition: grpc.ServiceDefinition<IAccessListService> = {
    getAccessLists: {
        path: "/teleport.accesslist.v1.AccessListService/GetAccessLists",
        originalName: "GetAccessLists",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => GetAccessListsResponse.fromBinary(bytes),
        requestDeserialize: bytes => GetAccessListsRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(GetAccessListsResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(GetAccessListsRequest.toBinary(value))
    },
    listAccessLists: {
        path: "/teleport.accesslist.v1.AccessListService/ListAccessLists",
        originalName: "ListAccessLists",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => ListAccessListsResponse.fromBinary(bytes),
        requestDeserialize: bytes => ListAccessListsRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(ListAccessListsResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(ListAccessListsRequest.toBinary(value))
    },
    getAccessList: {
        path: "/teleport.accesslist.v1.AccessListService/GetAccessList",
        originalName: "GetAccessList",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => AccessList.fromBinary(bytes),
        requestDeserialize: bytes => GetAccessListRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(AccessList.toBinary(value)),
        requestSerialize: value => Buffer.from(GetAccessListRequest.toBinary(value))
    },
    upsertAccessList: {
        path: "/teleport.accesslist.v1.AccessListService/UpsertAccessList",
        originalName: "UpsertAccessList",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => AccessList.fromBinary(bytes),
        requestDeserialize: bytes => UpsertAccessListRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(AccessList.toBinary(value)),
        requestSerialize: value => Buffer.from(UpsertAccessListRequest.toBinary(value))
    },
    updateAccessList: {
        path: "/teleport.accesslist.v1.AccessListService/UpdateAccessList",
        originalName: "UpdateAccessList",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => AccessList.fromBinary(bytes),
        requestDeserialize: bytes => UpdateAccessListRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(AccessList.toBinary(value)),
        requestSerialize: value => Buffer.from(UpdateAccessListRequest.toBinary(value))
    },
    deleteAccessList: {
        path: "/teleport.accesslist.v1.AccessListService/DeleteAccessList",
        originalName: "DeleteAccessList",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => Empty.fromBinary(bytes),
        requestDeserialize: bytes => DeleteAccessListRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(Empty.toBinary(value)),
        requestSerialize: value => Buffer.from(DeleteAccessListRequest.toBinary(value))
    },
    deleteAllAccessLists: {
        path: "/teleport.accesslist.v1.AccessListService/DeleteAllAccessLists",
        originalName: "DeleteAllAccessLists",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => Empty.fromBinary(bytes),
        requestDeserialize: bytes => DeleteAllAccessListsRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(Empty.toBinary(value)),
        requestSerialize: value => Buffer.from(DeleteAllAccessListsRequest.toBinary(value))
    },
    getAccessListsToReview: {
        path: "/teleport.accesslist.v1.AccessListService/GetAccessListsToReview",
        originalName: "GetAccessListsToReview",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => GetAccessListsToReviewResponse.fromBinary(bytes),
        requestDeserialize: bytes => GetAccessListsToReviewRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(GetAccessListsToReviewResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(GetAccessListsToReviewRequest.toBinary(value))
    },
    countAccessListMembers: {
        path: "/teleport.accesslist.v1.AccessListService/CountAccessListMembers",
        originalName: "CountAccessListMembers",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => CountAccessListMembersResponse.fromBinary(bytes),
        requestDeserialize: bytes => CountAccessListMembersRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(CountAccessListMembersResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(CountAccessListMembersRequest.toBinary(value))
    },
    listAccessListMembers: {
        path: "/teleport.accesslist.v1.AccessListService/ListAccessListMembers",
        originalName: "ListAccessListMembers",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => ListAccessListMembersResponse.fromBinary(bytes),
        requestDeserialize: bytes => ListAccessListMembersRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(ListAccessListMembersResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(ListAccessListMembersRequest.toBinary(value))
    },
    listAllAccessListMembers: {
        path: "/teleport.accesslist.v1.AccessListService/ListAllAccessListMembers",
        originalName: "ListAllAccessListMembers",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => ListAllAccessListMembersResponse.fromBinary(bytes),
        requestDeserialize: bytes => ListAllAccessListMembersRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(ListAllAccessListMembersResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(ListAllAccessListMembersRequest.toBinary(value))
    },
    getAccessListMember: {
        path: "/teleport.accesslist.v1.AccessListService/GetAccessListMember",
        originalName: "GetAccessListMember",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => Member.fromBinary(bytes),
        requestDeserialize: bytes => GetAccessListMemberRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(Member.toBinary(value)),
        requestSerialize: value => Buffer.from(GetAccessListMemberRequest.toBinary(value))
    },
    getStaticAccessListMember: {
        path: "/teleport.accesslist.v1.AccessListService/GetStaticAccessListMember",
        originalName: "GetStaticAccessListMember",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => GetStaticAccessListMemberResponse.fromBinary(bytes),
        requestDeserialize: bytes => GetStaticAccessListMemberRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(GetStaticAccessListMemberResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(GetStaticAccessListMemberRequest.toBinary(value))
    },
    getAccessListOwners: {
        path: "/teleport.accesslist.v1.AccessListService/GetAccessListOwners",
        originalName: "GetAccessListOwners",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => GetAccessListOwnersResponse.fromBinary(bytes),
        requestDeserialize: bytes => GetAccessListOwnersRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(GetAccessListOwnersResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(GetAccessListOwnersRequest.toBinary(value))
    },
    upsertAccessListMember: {
        path: "/teleport.accesslist.v1.AccessListService/UpsertAccessListMember",
        originalName: "UpsertAccessListMember",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => Member.fromBinary(bytes),
        requestDeserialize: bytes => UpsertAccessListMemberRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(Member.toBinary(value)),
        requestSerialize: value => Buffer.from(UpsertAccessListMemberRequest.toBinary(value))
    },
    upsertStaticAccessListMember: {
        path: "/teleport.accesslist.v1.AccessListService/UpsertStaticAccessListMember",
        originalName: "UpsertStaticAccessListMember",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => UpsertStaticAccessListMemberResponse.fromBinary(bytes),
        requestDeserialize: bytes => UpsertStaticAccessListMemberRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(UpsertStaticAccessListMemberResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(UpsertStaticAccessListMemberRequest.toBinary(value))
    },
    updateAccessListMember: {
        path: "/teleport.accesslist.v1.AccessListService/UpdateAccessListMember",
        originalName: "UpdateAccessListMember",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => Member.fromBinary(bytes),
        requestDeserialize: bytes => UpdateAccessListMemberRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(Member.toBinary(value)),
        requestSerialize: value => Buffer.from(UpdateAccessListMemberRequest.toBinary(value))
    },
    deleteAccessListMember: {
        path: "/teleport.accesslist.v1.AccessListService/DeleteAccessListMember",
        originalName: "DeleteAccessListMember",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => Empty.fromBinary(bytes),
        requestDeserialize: bytes => DeleteAccessListMemberRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(Empty.toBinary(value)),
        requestSerialize: value => Buffer.from(DeleteAccessListMemberRequest.toBinary(value))
    },
    deleteStaticAccessListMember: {
        path: "/teleport.accesslist.v1.AccessListService/DeleteStaticAccessListMember",
        originalName: "DeleteStaticAccessListMember",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => DeleteStaticAccessListMemberResponse.fromBinary(bytes),
        requestDeserialize: bytes => DeleteStaticAccessListMemberRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(DeleteStaticAccessListMemberResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(DeleteStaticAccessListMemberRequest.toBinary(value))
    },
    deleteAllAccessListMembersForAccessList: {
        path: "/teleport.accesslist.v1.AccessListService/DeleteAllAccessListMembersForAccessList",
        originalName: "DeleteAllAccessListMembersForAccessList",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => Empty.fromBinary(bytes),
        requestDeserialize: bytes => DeleteAllAccessListMembersForAccessListRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(Empty.toBinary(value)),
        requestSerialize: value => Buffer.from(DeleteAllAccessListMembersForAccessListRequest.toBinary(value))
    },
    deleteAllAccessListMembers: {
        path: "/teleport.accesslist.v1.AccessListService/DeleteAllAccessListMembers",
        originalName: "DeleteAllAccessListMembers",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => Empty.fromBinary(bytes),
        requestDeserialize: bytes => DeleteAllAccessListMembersRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(Empty.toBinary(value)),
        requestSerialize: value => Buffer.from(DeleteAllAccessListMembersRequest.toBinary(value))
    },
    upsertAccessListWithMembers: {
        path: "/teleport.accesslist.v1.AccessListService/UpsertAccessListWithMembers",
        originalName: "UpsertAccessListWithMembers",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => UpsertAccessListWithMembersResponse.fromBinary(bytes),
        requestDeserialize: bytes => UpsertAccessListWithMembersRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(UpsertAccessListWithMembersResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(UpsertAccessListWithMembersRequest.toBinary(value))
    },
    listAccessListReviews: {
        path: "/teleport.accesslist.v1.AccessListService/ListAccessListReviews",
        originalName: "ListAccessListReviews",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => ListAccessListReviewsResponse.fromBinary(bytes),
        requestDeserialize: bytes => ListAccessListReviewsRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(ListAccessListReviewsResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(ListAccessListReviewsRequest.toBinary(value))
    },
    listAllAccessListReviews: {
        path: "/teleport.accesslist.v1.AccessListService/ListAllAccessListReviews",
        originalName: "ListAllAccessListReviews",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => ListAllAccessListReviewsResponse.fromBinary(bytes),
        requestDeserialize: bytes => ListAllAccessListReviewsRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(ListAllAccessListReviewsResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(ListAllAccessListReviewsRequest.toBinary(value))
    },
    createAccessListReview: {
        path: "/teleport.accesslist.v1.AccessListService/CreateAccessListReview",
        originalName: "CreateAccessListReview",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => CreateAccessListReviewResponse.fromBinary(bytes),
        requestDeserialize: bytes => CreateAccessListReviewRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(CreateAccessListReviewResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(CreateAccessListReviewRequest.toBinary(value))
    },
    deleteAccessListReview: {
        path: "/teleport.accesslist.v1.AccessListService/DeleteAccessListReview",
        originalName: "DeleteAccessListReview",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => Empty.fromBinary(bytes),
        requestDeserialize: bytes => DeleteAccessListReviewRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(Empty.toBinary(value)),
        requestSerialize: value => Buffer.from(DeleteAccessListReviewRequest.toBinary(value))
    },
    accessRequestPromote: {
        path: "/teleport.accesslist.v1.AccessListService/AccessRequestPromote",
        originalName: "AccessRequestPromote",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => AccessRequestPromoteResponse.fromBinary(bytes),
        requestDeserialize: bytes => AccessRequestPromoteRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(AccessRequestPromoteResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(AccessRequestPromoteRequest.toBinary(value))
    },
    getSuggestedAccessLists: {
        path: "/teleport.accesslist.v1.AccessListService/GetSuggestedAccessLists",
        originalName: "GetSuggestedAccessLists",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => GetSuggestedAccessListsResponse.fromBinary(bytes),
        requestDeserialize: bytes => GetSuggestedAccessListsRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(GetSuggestedAccessListsResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(GetSuggestedAccessListsRequest.toBinary(value))
    },
    getInheritedGrants: {
        path: "/teleport.accesslist.v1.AccessListService/GetInheritedGrants",
        originalName: "GetInheritedGrants",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => GetInheritedGrantsResponse.fromBinary(bytes),
        requestDeserialize: bytes => GetInheritedGrantsRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(GetInheritedGrantsResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(GetInheritedGrantsRequest.toBinary(value))
    }
};
