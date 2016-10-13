var app = app || {};

(function() {
    app.Grid = function(presentations) {
        return new Backgrid.Grid({
            columns: [{
                name: "title",
                cell: "string",
                label: "Title",
                sortable: true,
                editable: false
            }, {
                name: "creator.name",
                cell: "uri",
                label: "Creator",
                sortable: true,
                editable: false
            }, {
                name: "createdAt",
                label: "Created At",
                cell: Backgrid.Extension.MomentCell.extend({
                    modelFormat: "YYYY/M/D",
                    displayLang: "zh-tw",
                    displayFormat: "MMM DD, YYYY"
                }),
                sortable: true,
                editable: false
            }],

            collection: presentations
        });
    }

})(jQuery)
