import * as ledger from './ledger';
import * as ui from './ui';

export { ledger, ui };

export type MoneypennyAction = ledger.LedgerAction | ui.UiAction;
