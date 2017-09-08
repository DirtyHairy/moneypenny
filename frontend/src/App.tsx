import * as React from 'react';
import {Store} from 'redux';
import {Provider} from 'react-redux';
import {MuiThemeProvider, createMuiTheme} from 'material-ui/styles';
import {AppBar, Toolbar, Typography} from 'material-ui';

import State from './state/state';

export interface Props {
    store: Store<State>;
}

const theme = createMuiTheme();

const App = (props: Props) =>
    <Provider store={props.store}>
    <MuiThemeProvider theme={theme}>
        <AppBar>
            <Toolbar>
                <Typography type='title' color='inherit'>
                    Moneypenny
                </Typography>
            </Toolbar>
        </AppBar>
    </MuiThemeProvider>
    </Provider>;

export default App;
