type Filter = "all" | "active" | "completed";

type FilterBarProps = {
  value: Filter;
  onChange: (value: Filter) => void;
};

export default function FilterBar({ value, onChange }: FilterBarProps) {
  return (
    <div className="filter-bar">
      {(["all", "active", "completed"] as const).map((option) => (
        <button
          key={option}
          className={value === option ? "active" : ""}
          onClick={() => onChange(option)}
          type="button"
        >
          {option}
        </button>
      ))}
    </div>
  );
}