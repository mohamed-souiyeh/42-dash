import { useEffect, useState } from "react";
import GameCard from "./gameCard";
import { ADDR, PORT } from "../../../endpoints";
import type { FiltersType } from "./filterTypes";

function RenderGames({
    query, setCategories, setProviders, filters,
}: {
    query: string,
    setCategories: (value: []) => void,
    setProviders: (value: []) => void,
    filters: FiltersType,
}) {

    const [data, setData] = useState([]);
    const [meta, setMeta] = useState([])
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        let r = `http://${ADDR}:${PORT}/api/games`;

        if (query !== "") { r += `?name=${query}`; }

        setLoading(true);
        fetch(r)
            .then((res) => res.json())
            .then((res) => {
                setData(res['data']);
                setMeta(res['meta']);
                setLoading(false);
            });
    }, [query]);

    useEffect(() => {
        if (!data || data.length === 0) return;

        const categoriesList: any = [];
        const providersList: any = [];

        data.forEach((e) => {
            const category = e['category'];
            const provider = e['provider'];

            if (!categoriesList.includes(category)) {
                categoriesList.push(category);
            }

            if (!providersList.includes(provider)) {
                providersList.push(provider);
            }
        });

        setCategories(categoriesList);
        setProviders(providersList);

    }, [data]);

    if (loading || !data || !meta) return <p>Loading Games...</p>;

    return (
        <div className="flex flex-col items-center gap-4">
            {data
                .filter((e: any) =>
                    filters.providers.length === 0 ||
                    filters.providers.includes(e.provider)
                )
                .filter((e: any) =>
                    filters.categories.length === 0 ||
                    filters.categories.includes(e.category)
                )
                .map((e: any) => (
                    <GameCard
                        key={e.id}
                        name={e.name}
                        provider={e.provider}
                        category={e.category}
                    />
                ))}
        </div>
    );
}

export default RenderGames;