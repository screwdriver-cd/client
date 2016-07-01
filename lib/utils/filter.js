'use strict';

/**
 * [description]
 * @method
 * @param  {[type]} list    [description]
 * @param  {[type]} filters [description]
 * @return {[type]}         [description]
 */
module.exports = (list, filters) => {
    let newList = list;

    Object.keys(filters).forEach((filter) => {
        if (filters[filter]) {
            newList = newList.filter((item) => item[filter] === filters[filter]);
        }
    });

    return newList;
};
