import { } from 'material-ui/colors/grey';
import * as React from 'react';
import Button from 'material-ui/Button';
import {MuiThemeProvider, createMuiTheme} from 'material-ui/styles';
import {AppBar, Toolbar, Typography} from 'material-ui';

export interface Props {}

const theme = createMuiTheme();

const App = () =>
    <MuiThemeProvider theme={theme}>
        <AppBar>
            <Toolbar>
                <Typography type='title' color='inherit'>
                    Moneypenny
                </Typography>
            </Toolbar>
        </AppBar>
    </MuiThemeProvider>;

export default App;
