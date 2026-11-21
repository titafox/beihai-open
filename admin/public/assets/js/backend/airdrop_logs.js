define(['jquery', 'bootstrap', 'backend', 'table', 'form'], function ($, undefined, Backend, Table, Form) {

    var Controller = {
        index: function () {
            // 初始化表格参数配置
            Table.api.init({
                extend: {
                    index_url: 'airdrop_logs/index' + location.search,
                    add_url: 'airdrop_logs/add',
                    edit_url: 'airdrop_logs/edit',
                    del_url: 'airdrop_logs/del',
                    multi_url: 'airdrop_logs/multi',
                    import_url: 'airdrop_logs/import',
                    table: 'tbl_airdrops_log',
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
                        {field: 'account_id', title: __('Account_id')},
                        {field: 'airdrops_id', title: __('Airdrops_id')},
                        {field: 'p_id', title: __('P_id')},
                        {field: 'n_id', title: __('N_id')},
                        {field: 'operation_id', title: __('Operation_id'), operate: 'LIKE'},
                        {field: 'status', title: __('Status')},
                        {field: 'message', title: __('Message'), operate: 'LIKE'},
                        {field: 'create_time', title: __('Create_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},
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
