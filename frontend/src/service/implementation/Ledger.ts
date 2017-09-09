import {Store} from 'redux';

import {RequestStatus} from '../../model/RequestResult';
import State from '../../state/State';
import TransactionProvider from '../provider/Transaction';
import {replaceAllTransactions} from '../../action/ledger';
import LedgerServiceInterface from '../Ledger';

class LedgerService implements LedgerServiceInterface {

    constructor(
        private _transactionProvider: TransactionProvider
    ) {}

    setStore(store: Store<State>): void {
        this._store = store;
    }

    async start(): Promise<void> {
        const result = await this._transactionProvider.loadAll();

        if (result.status === RequestStatus.success) {
            await this._store.dispatch(replaceAllTransactions(result.payload));
        }
    }

    private _store: Store<State>;

}

export default LedgerService;
