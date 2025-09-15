import React, { useEffect, useState } from "react";
import { RemotePlugin } from "./types";
import { PluginGrid } from "./PluginGrid";

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
            <PluginGrid plugins={plugins} />
        </div>
    );
}
