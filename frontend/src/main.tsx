import * as React from 'react';
import {render} from 'react-dom';

import App from './component/App';

const Main = () => <App/>;

render(<Main/>, document.getElementById('react-attachpoint'));
