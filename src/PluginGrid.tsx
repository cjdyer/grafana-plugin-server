import React, { useState } from "react";
import { RemotePlugin } from "./types";
import { RemotePluginCard } from "./PluginCard";
import { css } from "@emotion/css";

const getStyles = () => ({
    grid: css({
        display: "grid",
        gridTemplateColumns: "repeat(auto-fill, minmax(320px, 1fr))",
        gap: 16,
        marginTop: 24,
    }),
    searchBar: css({
        margin: "16px 0",
        input: {
            padding: 8,
            width: "100%",
            maxWidth: 400,
            fontSize: "1rem",
            border: "1px solid #ccc",
            borderRadius: 8,
        },
    }),
});

interface PluginGridProps {
    plugins: RemotePlugin[];
}

export function PluginGrid({ plugins }: PluginGridProps) {
    const [query, setQuery] = useState("");

    const styles = getStyles();

    console.log(plugins);

    const filteredPlugins = plugins.filter((p) =>
        p.id.toLowerCase().includes(query.toLowerCase())
    );

    return (
        <div>
            <div className={styles.searchBar}>
                <input
                    type="text"
                    placeholder="Search plugins by name…"
                    value={query}
                    onChange={(e) => setQuery(e.target.value)}
                />
            </div>
            {filteredPlugins.length === 0 ? (
                <p>No plugins found</p>
            ) : (
                <div className={styles.grid}>
                    {filteredPlugins.map((p) => (
                        <RemotePluginCard plugin={p} key={p.id} />
                    ))}
                </div>
            )}
        </div>
    );
}
