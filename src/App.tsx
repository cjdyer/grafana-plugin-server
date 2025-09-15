import React, { useEffect, useState } from "react";
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
});

export default function App() {
    const [plugins, setPlugins] = useState<RemotePlugin[]>([]);
    const [loading, setLoading] = useState(true);

    const styles = getStyles();

    useEffect(() => {
        async function fetchPlugins() {
            try {
                const res = await fetch("/api/plugins");
                const { items: remotePlugins }: { items: RemotePlugin[] } =
                    await res.json();
                console.log(remotePlugins);
                setPlugins(remotePlugins);
            } catch (err) {
                console.error("Failed to fetch plugins", err);
            } finally {
                setLoading(false);
            }
        }
        fetchPlugins();
    }, []);

    if (loading) {
        return <div>Loading plugins…</div>;
    }

    return (
        <div>
            <h1>Private Grafana Plugin Store</h1>
            {plugins.length === 0 ? (
                <p>No plugins found</p>
            ) : (
                <div className={styles.grid}>
                    {plugins.map((p) => (
                        <RemotePluginCard plugin={p} key={p.id} />
                    ))}
                </div>
            )}
        </div>
    );
}
