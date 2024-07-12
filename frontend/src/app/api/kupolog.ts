/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export type ProblemProblem = object;

export interface ServerAccountRequest {
  access_token?: string;
  expires_at?: number;
  provider?: string;
  provider_account_id?: string;
}

export interface ServerLoginResponse {
  token?: ServerLoginTokenResponse;
  user?: ServerLoginUserResponse;
}

export interface ServerLoginTokenResponse {
  access_token?: string;
  refresh_token?: string;
}

export interface ServerLoginUserResponse {
  createdAt?: string;
  email?: string;
  image?: string;
  name?: string;
  userID?: string;
}

export interface ServerServerInfo {
  environment?: string;
  version?: string;
}

/** Response for the health check */
export interface ServerServerStatus {
  /** Status is the health status of the service */
  status?: string;
  /** SystemInfo contains information about the system */
  system_info?: ServerServerInfo;
}

export interface ServerSharedTaskResponse {
  dailies?: ServerTaskResponse[];
  weeklies?: ServerTaskResponse[];
}

export interface ServerTaskResponse {
  taskID?: string;
  title?: string;
}

import type { AxiosInstance, AxiosRequestConfig, AxiosResponse, HeadersDefaults, ResponseType } from "axios";
import axios from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
  Text = "text/plain",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "/v1" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  protected mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    const method = params1.method || (params2 && params2.method);

    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...((method && this.instance.defaults.headers[method.toLowerCase() as keyof HeadersDefaults]) || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  protected stringifyFormItem(formItem: unknown) {
    if (typeof formItem === "object" && formItem !== null) {
      return JSON.stringify(formItem);
    } else {
      return `${formItem}`;
    }
  }

  protected createFormData(input: Record<string, unknown>): FormData {
    if (input instanceof FormData) {
      return input;
    }
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      const propertyContent: any[] = property instanceof Array ? property : [property];

      for (const formItem of propertyContent) {
        const isFileType = formItem instanceof Blob || formItem instanceof File;
        formData.append(key, isFileType ? formItem : this.stringifyFormItem(formItem));
      }

      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = format || this.format || undefined;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      body = this.createFormData(body as Record<string, unknown>);
    }

    if (type === ContentType.Text && body && body !== null && typeof body !== "string") {
      body = JSON.stringify(body);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(requestParams.headers || {}),
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title Swagger Kupolog API
 * @version 1.0
 * @license MIT (https://opensource.org/license/mit)
 * @termsOfService https://api.kupolog.com/terms
 * @baseUrl /v1
 * @externalDocs https://swagger.io/resources/open-api/
 * @contact nishojib <nishojib@kupolog.com> (https://api.kupolog.com/support)
 *
 * This is an API for the Kupolog app.
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  auth = {
    /**
     * @description takes a google or discord account request verifies the account and returns a token
     *
     * @tags auth
     * @name LoginCreate
     * @summary login
     * @request POST:/auth/login
     */
    loginCreate: (request: ServerAccountRequest, params: RequestParams = {}) =>
      this.request<
        ServerLoginResponse,
        {
          detail?: string;
          status?: number;
          title?: string;
          type?: string;
        }
      >({
        path: `/auth/login`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description refreshes the access token for the user
     *
     * @tags auth
     * @name RefreshCreate
     * @summary refresh token
     * @request POST:/auth/refresh
     * @secure
     */
    refreshCreate: (params: RequestParams = {}) =>
      this.request<
        {
          access_token?: string;
        },
        {
          detail?: string;
          status?: number;
          title?: string;
          type?: string;
        }
      >({
        path: `/auth/refresh`,
        method: "POST",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description revokes the refresh token for the user
     *
     * @tags auth
     * @name RevokeCreate
     * @summary revoke token
     * @request POST:/auth/revoke
     * @secure
     */
    revokeCreate: (params: RequestParams = {}) =>
      this.request<
        void,
        {
          detail?: string;
          status?: number;
          title?: string;
          type?: string;
        }
      >({
        path: `/auth/revoke`,
        method: "POST",
        secure: true,
        type: ContentType.Json,
        ...params,
      }),
  };
  health = {
    /**
     * @description Checks the health of the service
     *
     * @tags health
     * @name HealthList
     * @summary Health check
     * @request GET:/health
     */
    healthList: (params: RequestParams = {}) =>
      this.request<ServerServerStatus, ProblemProblem>({
        path: `/health`,
        method: "GET",
        format: "json",
        ...params,
      }),
  };
  tasks = {
    /**
     * @description Get the shared tasks
     *
     * @tags tasks
     * @name SharedList
     * @summary Shared tasks
     * @request GET:/tasks/shared
     */
    sharedList: (
      query: {
        /** Kind of tasks to return */
        kind: string;
      },
      params: RequestParams = {},
    ) =>
      this.request<
        ServerSharedTaskResponse,
        {
          detail?: string;
          status?: number;
          title?: string;
          type?: string;
        }
      >({
        path: `/tasks/shared`,
        method: "GET",
        query: query,
        format: "json",
        ...params,
      }),
  };
}
