import React, { useEffect, useState } from "react";
import { RemotePlugin } from "./types";
import { PluginGrid } from "./PluginGrid";
import { css } from "@emotion/css";

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
    uploadBox: css({
        margin: "16px 0 32px",
        padding: 16,
        border: "2px dashed #ccc",
        borderRadius: 8,
        textAlign: "center",
        background: "#fafafa",
        'input[type="file"]': {
            marginTop: 12,
        },
    }),
});

export default function App() {
    const [plugins, setPlugins] = useState<RemotePlugin[]>([]);
    const [loading, setLoading] = useState(true);
    const [uploading, setUploading] = useState(false);

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

    async function handleUpload(e: React.ChangeEvent<HTMLInputElement>) {
        if (!e.target.files || e.target.files.length === 0) return;

        const file = e.target.files[0];
        const text = await file.text();

        try {
            const json = JSON.parse(text);
            setUploading(true);
            const res = await fetch("/api/plugins", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(json),
            });

            if (!res.ok) {
                const msg = await res.text();
                alert("Upload failed: " + msg);
            } else {
                alert("Plugin uploaded successfully");
                fetchPlugins();
            }
        } catch (err) {
            alert("Invalid plugin.json file");
        } finally {
            setUploading(false);
            e.target.value = "";
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

            <div className={styles.uploadBox}>
                <p>
                    Upload a <code>plugin.json</code> file to register a new
                    plugin
                </p>
                <input
                    type="file"
                    accept="application/json"
                    onChange={handleUpload}
                    disabled={uploading}
                />
            </div>

            <PluginGrid plugins={plugins} />
        </div>
    );
}
