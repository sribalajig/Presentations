var app = app || {};

(function() {
    app.Presentations = Backbone.PageableCollection.extend({

        url: "http://localhost:8092/v1/presentations",
        searchText: "",
        state: {
            pageSize: 10,
            sortKey: "createdAt",
            order: 1
        },

        queryParams: {
            totalPages: null,
            totalRecords: null,
            sortKey: "sort",
            pageSize: "numitems",
            currentPage: "index",
            sortKey: "sortby",
            order: "direction",
            directions: { "0": "asc", "1": "desc" },
            paginate: true,
            sort: true,
            filter: 'title_'
        },

        parseState: function(resp, queryParams, state, options) {
            return {
                totalRecords: resp.totalRecords
            };
        },

        parseRecords: function(resp, options) {
            return resp.results;
        },

        search: function(searchParam) {
            this.state.currentPage = 1
            this.queryParams["filter"] = "title_" + searchParam

            this.fetch();
        }
    });
})();
