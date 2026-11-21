define(['jquery', 'bootstrap', 'backend', 'table', 'form'], function ($, undefined, Backend, Table, Form) {

    var Controller = {
        index: function () {
            // 初始化表格参数配置
            Table.api.init({
                extend: {
                    index_url: 'tbl/works/index' + location.search,
                    add_url: 'tbl/works/add',
                    edit_url: 'tbl/works/edit',
                    del_url: 'tbl/works/del',
                    multi_url: 'tbl/works/multi',
                    import_url: 'tbl/works/import',
                    table: 'tbl_works',
                }
            });

            var table = $("#table");

            // 初始化表格
            table.bootstrapTable({
                url: $.fn.bootstrapTable.defaults.extend.index_url,
                pk: 'id',
                sortName: 'id',
                fixedColumns: true,
                fixedRightNumber: 1,
                columns: [
                    [
                        {checkbox: true},
                        {field: 'id', title: __('Id')},
                        {field: 'ips.nickname', title: __('Ips.nickname'), operate: 'LIKE'},
                        {field: 'typ', title: __('Typ'), searchList: {"1":__('Typ 1'),"2":__('Typ 2'),"3":__('Typ 3')}, formatter: Table.api.formatter.normal},
                        {field: 'is_synthetic_material', title: __('Is_synthetic_material'), searchList: {"0":__('Is_synthetic_material 0'),"1":__('Is_synthetic_material 1')}, formatter: Table.api.formatter.normal},
                        {field: 'name', title: __('Name'), operate: 'LIKE'},
                        {field: 'image', title: __('Image'), operate: false, events: Table.api.events.image, formatter: Table.api.formatter.image},
                        {field: 'price', title: __('Price'), operate:'BETWEEN'},
                        {field: 'stock', title: __('Stock')},
                        {field: 'reserved', title: __('Reserved')},
                        {field: 'stock_sell', title: __('Stock_sell')},
                        {field: 'advance_hour', title: __('Advance_hour')},
                        {field: 'pub_time', title: __('Pub_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},
                        {field: 'sell_time', title: __('Sell_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},
                        {field: 'sold_out', title: __('Sold_out'), searchList: {"0":__('Sold_out 0'),"1":__('Sold_out 1')}, formatter: Table.api.formatter.normal},
                        {field: 'status', title: __('Status'), searchList: {"0":__('Status 0'),"1":__('Status 1')}, formatter: Table.api.formatter.status},
                       // {field: 'create_time', title: __('Create_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},
                        {field: 'update_time', title: __('Update_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},


                        {field: 'operate', title: __('Operate'), table: table, events: Table.api.events.operate, formatter: Table.api.formatter.operate}
                    ]
                ]
            });

            // 为表格绑定事件
            Table.api.bindevent(table);
        },
        add: function () {
            Controller.api.bindevent();
        },
        edit: function () {
            Controller.api.bindevent();
        },
        api: {
            bindevent: function () {
                Form.api.bindevent($("form[role=form]"));
            }
        }
    };
    return Controller;
});
