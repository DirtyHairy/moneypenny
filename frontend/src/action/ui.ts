import {Action} from 'redux';

export enum Actions {
    setLoading = 'ui/set-loading'
}

export type UiAction = SetLoadingAction;

export interface SetLoadingAction extends Action {
    type: Actions.setLoading;
    loading: boolean;
}

export function setLoading(loading: boolean): SetLoadingAction {
    return {
        type: Actions.setLoading,
        loading
    };
}
