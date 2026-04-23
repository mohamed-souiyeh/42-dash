import { type ReactNode } from "react";
import type { FiltersType } from "./filterTypes";

function Filters({
    categories,
    providers,
    filters,
    setFilters,
}: {
    categories: string[];
    providers: string[];
    filters: FiltersType;
    setFilters: (value: FiltersType) => void;
}) {

    // This function updates the state of the active filters
    function toggleValue(filterType: keyof FiltersType, value: string) {
        setFilters({
            ...filters,
            [filterType]: filters[filterType].includes(value)
                ? filters[filterType].filter((v) => v !== value)
                : [...filters[filterType], value],
        });
    }

    // This helper component includes the header, including the filter it self as children
    function Section({ header, children, }: { header: string; children: ReactNode; }) {
        return (
            <div className="flex flex-col gap-3">
                <h2 className="font-semibold text-xl">| {header}</h2>
                <div className="flex flex-wrap gap-2">
                    {children}
                </div>
            </div>
        );
    }

    // Filter chip component
    function Chip({
        active,
        label,
        onClick,
        color = "green",
    }: {
        label: string;
        active: boolean;
        onClick: () => void;
        color?: "blue" | "green";
    }) {
        const colorMap = {
            blue: active
                ? "bg-blue-500 text-white"
                : "bg-blue-100 text-blue-700",
            green: active
                ? "bg-green-500 text-white"
                : "bg-green-100 text-green-700",
        };

        const classTags =
            "px-3 py-2 rounded-full text-sm transition " + colorMap[color];

        return (
            <button
                className={classTags}
                type="button"
                onClick={onClick}
                aria-pressed={active}
            >
                {label.toUpperCase()}
            </button>
        );
    }

    // This component features the filtering chips with a <select> object include it
    // They work on the same state managment, I'm doing this because playwright needs
    // to be happy, and I wan't to be happy too with my chips filter. 🍟
    function MultiSelectWithChips({
        options,
        filterType,
        testid,
        color,
    }: {
        options: string[];
        filterType: keyof FiltersType;
        testid: string;
        color: "blue" | "green";
    }) {
        return (
            <div className="flex flex-col gap-2">

                {/* Playwright is happy with this */}
                <select data-testid={`${testid}`}
                    // We make it invisible but its completly optional
                    // both systems work separatly and together but I
                    // decide I prefer chips. I already tried to use
                    // the select and option tags and change them visualy
                    // but the results doesn't convince me.
                    className="sr-only"
                    value=""
                    onChange={(e) => {
                        const value = e.target.value;
                        if (value) toggleValue(filterType, value);
                    }}
                >
                    <option value="">Add filter...</option>
                    {options.map((opt) => (
                        <option key={opt} value={opt}>
                            {opt}
                        </option>
                    ))}
                </select>

                {/* Now is Ismael turn to enjoy */}
                <div className="flex flex-wrap gap-3 max-w-md">
                    {options.length === 0 ? (
                        <p className="text-gray-500 text-sm">No available options at this moment.</p>
                    ) : (
                        options.map((opt) => (
                            <Chip
                                key={opt}
                                label={opt}
                                active={filters[filterType].includes(opt)}
                                onClick={() => toggleValue(filterType, opt)}
                                color={color}
                            />
                        ))
                    )}
                </div>
            </div>
        );
    }

    return (
        <div
            className="
                max-w-sm
                md:max-w-md
                xl:max-w-xl
                flex flex-col gap-12"
        >
            <Section header="Categories">
                <MultiSelectWithChips
                    options={categories}
                    filterType="categories"
                    testid="category-filter"
                    color="blue"
                />
            </Section>

            <Section header="Providers">
                <MultiSelectWithChips
                    options={providers}
                    filterType="providers"
                    testid="provider-filter"
                    color="green"
                />
            </Section>

        </div>
    );
}

export default Filters;