import { css } from "@emotion/css";
import { useState } from "react";

const getStyles = () => ({
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

interface Props {}

export function Upload({}: Props) {
    const [uploading, setUploading] = useState(false);

    const styles = getStyles();

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
                // fetchPlugins();
            }
        } catch (err) {
            alert("Invalid plugin.json file");
        } finally {
            setUploading(false);
            e.target.value = "";
        }
    }

    return (
        <div className={styles.uploadBox}>
            <p>
                Upload a <code>plugin.json</code> file to register a new plugin
            </p>
            <input
                type="file"
                accept="application/json"
                onChange={handleUpload}
                disabled={uploading}
            />
        </div>
    );
}
