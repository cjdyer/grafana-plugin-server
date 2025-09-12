import React, { useEffect, useState } from "react";
import { RemotePlugin } from "./types";

export default function App() {
    const [plugins, setPlugins] = useState<RemotePlugin[]>([]);
    const [loading, setLoading] = useState(true);

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
                <ul>
                    {plugins.map((p) => (
                        <li key={p.id}>
                            <strong>{p.id}</strong> ({p.name})<br />
                            Downloads: {p.downloads}
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
}
