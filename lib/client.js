'use strict';

const config = require('configrc').make('sd');
const Swagger = require('swagger-client');
const url = require('url');
const convert = require('./utils/convert');

/**
 * [getClient description]
 * @method getClient
 * @return {[type]}  [description]
 */
module.exports = () => {
    const token = config.get('token');
    const apiUrl = config.get('url');
    const auth = new Swagger.ApiKeyAuthorization(
        'Authorization',
        `Bearer ${token}`,
        'header'
    );

    if (!apiUrl) {
        return Promise.reject(new Error('URL missing, please set via `config url <api-url>`'));
    }

    if (!token) {
        return Promise.reject(new Error('Token missing, please login via `login`'));
    }

    return new Swagger({
        url: url.resolve(apiUrl, '/v3/swagger.json'),
        authorizations: { auth },
        usePromise: true
    }).catch((err) => {
        const message = convert(err);

        return Promise.reject(new Error(`Unable to access API, did you run setup?\n${message}`));
    });
};
