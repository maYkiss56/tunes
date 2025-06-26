interface Tab<T extends string> {
  id: T;
  label: string;
}

interface TabsProps<T extends string> {
  tabs: Tab<T>[];
  activeTab: T;
  onChange: (tabId: T) => void;
}

export const Tabs = <T extends string>({
  tabs,
  activeTab,
  onChange,
}: TabsProps<T>) => {
  return (
    <div className="flex border-b border-gray-800">
      {tabs.map((tab) => (
        <button
          key={tab.id}
          onClick={() => onChange(tab.id)}
          className={`px-4 py-2 font-medium text-sm relative cursor-pointer ${
            activeTab === tab.id
              ? "text-purple-500"
              : "text-gray-400 hover:text-white"
          }`}
        >
          {tab.label}
          {activeTab === tab.id && (
            <div className="absolute bottom-0 left-0 right-0 h-0.5 bg-purple-500 rounded-t" />
          )}
        </button>
      ))}
    </div>
  );
};
