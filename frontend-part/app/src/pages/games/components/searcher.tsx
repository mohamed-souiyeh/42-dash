function Search({ query, setQuery }: { query: string, setQuery: (value: string) => void }) {
    return (
        <div className="min-w-sm md:min-w-md xl:min-w-xl flex flex-col gap-6 mb-6">
            <h1 className="text-3xl whitespace-nowrap">Game List</h1>
            <div className="flex-1">
                <input
                    data-testid="search-input"
                    name="search"
                    className="w-full max-w-md max-h-12 border-2 border-gray-300 rounded-2xl px-4 py-2 
                        focus:outline-none focus:ring-2 focus:ring-green-200"
                    type="text"
                    placeholder="Search by name..."
                    value={query}
                    onChange={(e) => setQuery(e.target.value)}
                />
            </div>
        </div>
    );
}

export default Search;