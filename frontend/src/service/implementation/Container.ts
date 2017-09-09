import {Store} from 'redux';
import {Event} from 'microevent.ts';

import State from '../../state/State';
import ContainerInterface from '../Container';

import ApiRouter from './ApiRouter';
import LedgerService from './Ledger';
import TransactionProvider from '../provider/implementation/Transaction';

class Container implements ContainerInterface {

    setStore(store: Store<State>): this {
        if (this._store) {
            throw new Error('store already injected');
        }

        this._store = store;

        this._onStoreConfigured.dispatch(store);

        return this;
    }

    getApiRouter(): ApiRouter {
        return this._getOrCreateSingletonService('api-router', () => new ApiRouter());
    }

    getLedgerService(): LedgerService {
        return this._getOrCreateSingletonService(
            'ledger-service',
            () => new LedgerService(this.getTransactionProvider())
        );
    }

    getTransactionProvider(): TransactionProvider {
        return this._getOrCreateSingletonService(
            'transaction-provider',
            () => new TransactionProvider(this.getApiRouter())
        );
    }

    private _getOrCreateSingletonService(
        key: string,
        factory: () => any
    ): any {
        if (!this._services.has(key)) {
            const service = factory();

            this._services.set(key, service);

            if (service.setStore) {
                this._store ?
                    service.setStore(this._store) :
                    this._onStoreConfigured.addHandler(store => service.setStore(store));
            }
        }

        return this._services.get(key);
    }

    private _services = new Map<string, any>();

    private _onStoreConfigured = new Event<Store<State>>();

    private _store: Store<State> = null;

}

export default Container;
