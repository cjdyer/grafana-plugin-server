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

        if (!file.name.endsWith(".tar") && !file.name.endsWith(".tar.gz")) {
            alert("Please upload a .tar or .tar.gz plugin file");
            e.target.value = "";
            return;
        }

        const formData = new FormData();
        formData.append("plugin", file);

        try {
            setUploading(true);
            const res = await fetch("/api/plugins", {
                method: "POST",
                body: formData,
            });

            if (!res.ok) {
                const msg = await res.text();
                alert("Upload failed: " + msg);
            } else {
                alert("Plugin uploaded successfully");
                // fetchPlugins();
            }
        } catch (err) {
            alert("Upload failed: " + (err as Error).message);
        } finally {
            setUploading(false);
            e.target.value = "";
        }
    }

    return (
        <div className={styles.uploadBox}>
            <p>
                Upload a Grafana plugin as <code>.tar</code> or{" "}
                <code>.tar.gz</code>
            </p>
            <input
                type="file"
                accept=".tar,.tar.gz"
                onChange={handleUpload}
                disabled={uploading}
            />
        </div>
    );
}
