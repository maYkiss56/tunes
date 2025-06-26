import {
  type FC,
  useState,
  useRef,
  useEffect,
  type KeyboardEvent,
  type ChangeEvent,
} from "react";
import { ChevronDownIcon } from "./icons";

interface SelectOption {
  value: string;
  label: string;
}

interface SelectProps {
  options: SelectOption[];
  value?: string;
  onChange?: (e: ChangeEvent<HTMLSelectElement>) => void;
  className?: string;
  placeholder?: string;
  disabled?: boolean;
}

export const Select: FC<SelectProps> = ({
  options,
  value,
  onChange,
  className = "",
  placeholder = "Select...",
  disabled = false,
}) => {
  const [isOpen, setIsOpen] = useState(false);
  const [selectedIndex, setSelectedIndex] = useState<number | null>(null);
  const selectRef = useRef<HTMLDivElement>(null);
  const optionsRef = useRef<HTMLUListElement>(null);
  const buttonRef = useRef<HTMLButtonElement>(null);

  const [buttonWidth, setButtonWidth] = useState("auto");

  useEffect(() => {
    if (buttonRef.current) {
      setButtonWidth(`${buttonRef.current.offsetWidth}px`);
    }
  }, []);

  useEffect(() => {
    if (value) {
      const index = options.findIndex((option) => option.value === value);
      setSelectedIndex(index >= 0 ? index : null);
    }
  }, [value, options]);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        selectRef.current &&
        !selectRef.current.contains(event.target as Node)
      ) {
        setIsOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const handleKeyDown = (e: KeyboardEvent<HTMLDivElement>) => {
    if (!isOpen) {
      if (
        e.key === "Enter" ||
        e.key === " " ||
        e.key === "ArrowDown" ||
        e.key === "ArrowUp"
      ) {
        e.preventDefault();
        setIsOpen(true);
      }
      return;
    }

    switch (e.key) {
      case "Escape":
        e.preventDefault();
        setIsOpen(false);
        break;
      case "ArrowDown":
        e.preventDefault();
        setSelectedIndex((prev) => {
          const next =
            prev === null ? 0 : Math.min(prev + 1, options.length - 1);
          scrollToOption(next);
          return next;
        });
        break;
      case "ArrowUp":
        e.preventDefault();
        setSelectedIndex((prev) => {
          const next =
            prev === null ? options.length - 1 : Math.max(prev - 1, 0);
          scrollToOption(next);
          return next;
        });
        break;
      case "Enter":
      case " ":
        e.preventDefault();
        if (selectedIndex !== null) {
          handleSelect(options[selectedIndex].value);
          setIsOpen(false);
        }
        break;
    }
  };

  const scrollToOption = (index: number) => {
    if (optionsRef.current && index >= 0) {
      const optionElement = optionsRef.current.children[index] as HTMLElement;
      optionElement?.scrollIntoView({ block: "nearest" });
    }
  };

  const handleSelect = (value: string) => {
    if (onChange) {
      const fakeEvent = {
        target: {
          value: value,
          name: "",
          type: "select",
        },
        currentTarget: {
          value: value,
        },
        preventDefault: () => {},
        stopPropagation: () => {},
        nativeEvent: new Event("change"),
      } as unknown as ChangeEvent<HTMLSelectElement>;

      onChange(fakeEvent);
    }
    setIsOpen(false);
  };

  const selectedOption = selectedIndex !== null ? options[selectedIndex] : null;

  return (
    <div
      ref={selectRef}
      className={`relative ${className}`}
      onKeyDown={handleKeyDown}
      tabIndex={0}
    >
      <button
        ref={buttonRef}
        type="button"
        onClick={() => !disabled && setIsOpen(!isOpen)}
        disabled={disabled}
        className={`
          flex items-center justify-between w-full bg-gray-800 border 
          ${isOpen ? "border-purple-500 ring-2 ring-purple-500" : "border-gray-700"} 
          text-white rounded-lg pl-4 pr-10 py-2.5 focus:outline-none 
          transition-all duration-200 hover:border-gray-600
          ${disabled ? "opacity-50 cursor-not-allowed" : "cursor-pointer"}
          min-w-[190px]
        `}
        aria-haspopup="listbox"
        aria-expanded={isOpen}
      >
        <span className="truncate">
          {selectedOption ? selectedOption.label : placeholder}
        </span>
        <ChevronDownIcon
          className={`
            absolute right-3 w-5 h-5 text-gray-400 
            transition-transform duration-200 pointer-events-none
            ${isOpen ? "transform rotate-180" : ""}
          `}
        />
      </button>

      {/* Dropdown options */}
      {isOpen && (
        <ul
          ref={optionsRef}
          style={{ width: buttonWidth }}
          className={`
            absolute z-20 mt-1 max-h-60 overflow-auto
            bg-gray-800 border border-gray-700 rounded-lg shadow-lg
            focus:outline-none transition-opacity duration-200
            ${isOpen ? "opacity-100" : "opacity-0"}
            scrollbar-thin scrollbar-thumb-gray-700 scrollbar-track-gray-900`}
          role="listbox"
          aria-activedescendant={
            selectedIndex !== null ? `option-${selectedIndex}` : undefined
          }
        >
          {options.map((option, index) => (
            <li
              key={option.value}
              id={`option-${index}`}
              role="option"
              aria-selected={selectedIndex === index}
              className={`
                px-4 py-2.5 cursor-pointer transition-colors duration-150
                ${selectedIndex === index ? "bg-gray-700 text-purple-400" : "hover:bg-gray-700"}
                ${index === 0 ? "rounded-t-lg" : ""}
                ${index === options.length - 1 ? "rounded-b-lg" : ""}
                border-b ${index === options.length - 1 ? "border-transparent" : "border-gray-700"}
                truncate
              `}
              onClick={() => handleSelect(option.value)}
              onMouseEnter={() => setSelectedIndex(index)}
            >
              {option.label}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};
