import * as React from 'react';
import {render} from 'react-dom';

const Main = () =>
    <div>
        Hello world from react!
    </div>;

render(<Main/>, document.getElementById('react-attachpoint'));
