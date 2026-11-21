define(['jquery', 'bootstrap', 'backend', 'table', 'form'], function ($, undefined, Backend, Table, Form) {

    var Controller = {
        index: function () {
            // 初始化表格参数配置
            Table.api.init({
                extend: {
                    index_url: 'tbl/nfts/index' + location.search,
                    add_url: 'tbl/nfts/add',
                    edit_url: 'tbl/nfts/edit',
                    del_url: 'tbl/nfts/del',
                    multi_url: 'tbl/nfts/multi',
                    import_url: 'tbl/nfts/import',
                    table: 'tbl_nfts',
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
                        //{field: 'wid', title: __('Wid')},
                      {field: 'account.name', title: __('Account.name'), operate: 'LIKE'},
                      {field: 'account.phone', title: __('Account.phone'), operate: 'LIKE'},
                      {field: 'works.name', title: __('Works.name'), operate: 'LIKE'},
                        {field: 'work_no', title: __('Work_no')},
                        //{field: 'author_id', title: __('Author_id')},
                       // {field: 'owner_id', title: __('Owner_id')},
                        {field: 'transfer_num', title: __('Transfer_num')},
                        {field: 'price', title: __('Price'), operate:'BETWEEN'},
                       // {field: 'nft_id', title: __('Nft_id'), operate: 'LIKE'},
                       // {field: 'operation_id', title: __('Operation_id'), operate: 'LIKE'},
                        {field: 'chain_status', title: __('Chain_status'), searchList: {"0":__('Chain_status 0'),"1":__('Chain_status 1'),"2":__('Chain_status 2')}, formatter: Table.api.formatter.status},
                        {field: 'status', title: __('Status'), searchList: {"0":__('Status 0'),"1":__('Status 1'),"2":__('Status 2'),"3":__('Status 3')}, formatter: Table.api.formatter.status},
                        {field: 'update_time', title: __('Update_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},
                       // {field: 'mark', title: __('Mark'), operate: 'LIKE'},



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
