import { PencilIcon, TrashIcon } from "../../ui/icons";

export interface Column<T, K extends keyof T = keyof T> {
  key: K;
  header: string;
  render?: (value: T[K], item: T) => React.ReactNode;
}

interface DataTableProps<T extends { id: number }> {
  columns: Column<T>[];
  data: T[];
  onEdit: (item: T) => void;
  onDelete: (id: number) => void;
}

const DataTable = <T extends { id: number }>({
  columns,
  data,
  onEdit,
  onDelete,
}: DataTableProps<T>) => {
  return (
    <div className="overflow-x-auto rounded-lg border border-gray-200 shadow-sm">
      <table className="min-w-full divide-y divide-gray-200">
        <thead className="bg-gray-50">
          <tr>
            {columns.map((column) => (
              <th
                key={`header-${String(column.key)}`} // Уникальный ключ для заголовков
                scope="col"
                className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                {column.header}
              </th>
            ))}
            <th
              key="actions-header" // Фиксированный ключ для заголовка действий
              scope="col"
              className="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider"
            >
              Действия
            </th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {data.map((item) => {
            const rowKey = `row-${item.id}`; // Ключ для строки

            return (
              <tr key={rowKey} className="hover:bg-gray-50">
                {columns.map((column) => {
                  const cellKey = `${rowKey}-cell-${String(column.key)}`; // Ключ для ячейки

                  return (
                    <td
                      key={cellKey}
                      className="px-6 py-4 whitespace-nowrap text-sm text-gray-900"
                    >
                      {column.render
                        ? column.render(item[column.key], item)
                        : String(item[column.key])}
                    </td>
                  );
                })}
                <td
                  key={`${rowKey}-actions`} // Ключ для ячейки действий
                  className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2"
                >
                  <button
                    onClick={() => onEdit(item)}
                    className="text-purple-600 hover:text-purple-900"
                    aria-label="Редактировать"
                  >
                    <PencilIcon className="w-5 h-5" />
                  </button>
                  <button
                    onClick={() => onDelete(item.id)}
                    className="text-red-600 hover:text-red-900"
                    aria-label="Удалить"
                  >
                    <TrashIcon className="w-5 h-5" />
                  </button>
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
};
export default DataTable;
