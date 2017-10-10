import * as React from 'react';
import { render } from 'react-dom';
import { createStore } from 'redux';

import { isProduction } from './util';
import reducer from './reducer/root';
import ServiceContainerInterface from './service/Container';
import ServiceContainer from './service/implementation/Container';
import App from './App';

declare namespace window {
    export const devToolsExtension: any;
}

async function main() {
    const container: ServiceContainerInterface = new ServiceContainer();

    const store = createStore(
        reducer,
        isProduction || !window.devToolsExtension ? (x: any) => x : window.devToolsExtension()
    );

    container.setStore(store);

    await new Promise(r => render(<App store={store} />, document.getElementById('react-attachpoint'), r));

    await container.getLedgerService().start();
}

main();
