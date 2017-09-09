import {Store} from 'redux';

import State from '../state/State';

import ApiRouter from './ApiRouter';
import LedgerService from './Ledger';
import TransactionProvider from './provider/Transaction';

interface Container {

    setStore(store: Store<State>): this;

    getApiRouter(): ApiRouter;

    getLedgerService(): LedgerService;

    getTransactionProvider(): TransactionProvider;

}

export default Container;
