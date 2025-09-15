import { useEffect, useState } from "react";
import { RemotePlugin } from "./types";
import { PluginGrid } from "./PluginGrid";
import { css } from "@emotion/css";
import { Upload } from "./Upload";

const getStyles = () => ({
    appContainer: css({
        maxWidth: 1200,
        margin: "0 auto",
        padding: "32px 16px",
        fontFamily: "Arial, sans-serif",
    }),
    header: css({
        textAlign: "center",
        marginBottom: 32,
        h1: {
            fontSize: "2rem",
            margin: 0,
        },
        p: {
            marginTop: 8,
            color: "#555",
        },
    }),
    loading: css({
        textAlign: "center",
        padding: "64px 0",
        fontSize: "1.2rem",
    }),
});

export default function App() {
    const [plugins, setPlugins] = useState<RemotePlugin[]>([]);
    const [loading, setLoading] = useState(true);

    const styles = getStyles();

    useEffect(() => {
        fetchPlugins();
    }, []);

    async function fetchPlugins() {
        try {
            const res = await fetch("/api/plugins");
            const { items: remotePlugins }: { items: RemotePlugin[] } =
                await res.json();
            setPlugins(remotePlugins);
        } catch (err) {
            console.error("Failed to fetch plugins", err);
        } finally {
            setLoading(false);
        }
    }

    if (loading) {
        return <div className={styles.loading}>Loading plugins…</div>;
    }

    return (
        <div className={styles.appContainer}>
            <header className={styles.header}>
                <h1>Private Grafana Plugin Store</h1>
                <p>
                    Browse and manage your organization's custom Grafana plugins
                </p>
            </header>

            <Upload />
            <PluginGrid plugins={plugins} />
        </div>
    );
}
