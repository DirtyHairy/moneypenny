export const enum RequestStatus {
    success,
    error
}

export const enum ErrorReason {
    network,
    malformedResponse,
    domain
}

export interface RequestResponseSuccess<PayloadT = void> {
    status: RequestStatus.success;

    payload?: PayloadT;
}

export interface RequestResponseError<DomainErrorT = void> {
    status: RequestStatus.error;
    reason: ErrorReason;

    domainError?: DomainErrorT;
}

type RequestResult<PayloadT = void, DomainErrorT = void> =
    | RequestResponseSuccess<PayloadT>
    | RequestResponseError<DomainErrorT>;

namespace RequestResult {
    export function networkError(): RequestResponseError {
        return {
            status: RequestStatus.error,
            reason: ErrorReason.network
        };
    }

    export function malformedResponseError(): RequestResponseError {
        return {
            status: RequestStatus.error,
            reason: ErrorReason.malformedResponse
        };
    }

    export function success<PayloadT = void>(payload: PayloadT): RequestResponseSuccess<PayloadT> {
        return {
            status: RequestStatus.success,
            payload
        };
    }
}

export default RequestResult;
