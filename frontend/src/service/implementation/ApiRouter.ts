import ApiRouterInterface from '../ApiRouter';

class ApiRouter implements ApiRouterInterface {

    constructor(prefix = '') {
        this._prefix = prefix.replace(/\/+$/, '') + '/api/';
    }

    getAllTransactionsRoute(): string {
        return this._prefix + 'transactions';
    }

    private _prefix: string;

}

export default ApiRouter;
