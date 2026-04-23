import { useState } from 'react';
import Filters from './components/filters';
import Search from './components/searcher';
import RenderGames from './components/gameRendering';
import type { FiltersType } from './components/filterTypes';

export default function Games() {
    const [query, setQuery] = useState("");
    const [categories, setCategories] = useState([]);
    const [providers, setProviders] = useState([]);
    const [filters, setFilters] = useState<FiltersType>({
        providers: [],
        categories: [],
    });

    return (
        <div className="flex justify-center mx-12 my-12">
            <div className="flex flex-col items-center lg:items-start
                    lg:grid lg:grid-cols-2 gap-10 lg:gap-25 max-w-5xl w-full">

                <div className="flex flex-col gap-6">
                    <Search query={query} setQuery={setQuery} />

                    <Filters
                        filters={filters}
                        setFilters={setFilters}
                        categories={categories}
                        providers={providers}
                    />
                </div>

                <div className='px-6'>
                    <RenderGames
                        query={query}
                        setCategories={setCategories}
                        setProviders={setProviders}
                        filters={filters}
                    />
                </div>

            </div>
        </div>
    );
}