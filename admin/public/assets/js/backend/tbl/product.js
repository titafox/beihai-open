define(['jquery', 'bootstrap', 'backend', 'table', 'form'], function ($, undefined, Backend, Table, Form) {

    var Controller = {
        index: function () {
            // 初始化表格参数配置
            Table.api.init({
                extend: {
                    index_url: 'tbl/product/index' + location.search,
                    add_url: 'tbl/product/add',
                    edit_url: 'tbl/product/edit',
                    del_url: 'tbl/product/del',
                    multi_url: 'tbl/product/multi',
                    import_url: 'tbl/product/import',
                    table: 'tbl_product',
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
                        {field: 'account.phone', title: __('Author'), operate: 'LIKE'},
                        {field: 'owner.phone', title: __('Owner'),operate: 'LIKE'},
                       // {field: 'author_id', title: __('Author_id')},
                        //{field: 'classify', title: __('Classify')},
                       // {field: 'tx_id', title: __('Tx_id'), operate: 'LIKE'},
                        {field: 'token_id', title: __('Token_id'), operate: 'LIKE'},
                        {field: 'name', title: __('Name'), operate: 'LIKE'},
                        {field: 'image', title: __('Image'), operate: false, events: Table.api.events.image, formatter: Table.api.formatter.image},
                        {field: 'price', title: __('Price')},
                        {field: 'status', title: __('Status')},
                        {field: 'stock', title: __('Stock')},
                        {field: 'index', title: __('Index')},
                        {field: 'order_no', title: __('Order_no'), operate: 'LIKE'},
                        {field: 'create_time', title: __('Create_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},
                        {field: 'update_time', title: __('Update_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},
                        {field: 'del_time', title: __('Del_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},
                        {field: 'chain_status', title: __('Chain_status')},
                        {field: 'hash', title: __('Hash'), operate: 'LIKE'},
                        {field: 'pay_type', title: __('Pay_type')},
                        {field: 'c_name', title: __('C_name'), operate: 'LIKE'},
                        {field: 'sale_time', title: __('Sale_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},
                        {field: 'tinyint', title: __('Tinyint')},
                        {field: 'stock_sell', title: __('Stock_sell')},
                        {field: 'circulation', title: __('Circulation')},

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
