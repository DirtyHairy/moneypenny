import * as moment from 'moment';

export const isProduction = process.env['NODE_ENV'] === 'production';

export function parseRfc3339(date: string): Date {
    return new Date(moment(date).unix() * 1000);
}
