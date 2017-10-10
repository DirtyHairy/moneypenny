import { default as TransactionProviderInterface, LoadAllResult } from '../Transaction';
import { RequestResult, Transaction } from '../../../model';
import ApiRouter from '../../ApiRouter';
import { parseRfc3339 } from '../../../util';

class TransactionProvider implements TransactionProviderInterface {
    constructor(private _router: ApiRouter, private _fetch = fetch.bind(window)) {}

    async loadAll(): Promise<LoadAllResult> {
        let response: Response;

        try {
            response = await this._fetch(this._router.getAllTransactionsRoute());
        } catch (e) {
            return RequestResult.networkError();
        }

        if (response.status !== 200) {
            return RequestResult.malformedResponseError();
        }

        let rawTransactions: Array<Transaction>;

        try {
            rawTransactions = await response.json();
        } catch (e) {
            return RequestResult.malformedResponseError();
        }

        const transactions = rawTransactions.map(
            (t: any) =>
                ({
                    ...t,
                    transactionDate: parseRfc3339(t.transactionDate)
                } as Transaction)
        );

        return RequestResult.success(transactions);
    }
}

export default TransactionProvider;
