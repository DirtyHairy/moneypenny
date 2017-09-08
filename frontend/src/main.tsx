import * as React from 'react';
import {render} from 'react-dom';
import {createStore} from 'redux';

import {isProduction} from './util';
import reducer from './reducer/root';
import App from './App';

declare namespace window {
    export const devToolsExtension: any;
}

function main() {
    const store = createStore(
        reducer,
        (isProduction || !window.devToolsExtension) ? (x: any) => x : window.devToolsExtension()
    );

    render(<App store={store}/>,
    document.getElementById('react-attachpoint'));
}

main();
