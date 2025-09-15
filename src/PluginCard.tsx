import React from "react";
import {
    Card,
    CardContent,
    CardHeader,
    CardActions,
    Typography,
    Chip,
    Box,
    Button,
    Avatar,
    Tooltip,
    Stack,
    Divider,
} from "@mui/material";
import VerifiedIcon from "@mui/icons-material/Verified";
import DownloadIcon from "@mui/icons-material/Download";
import LinkIcon from "@mui/icons-material/Link";
import BusinessIcon from "@mui/icons-material/Business";
import { css } from "@emotion/css";

import { RemotePlugin } from "./types";

interface RemotePluginCardProps {
    plugin: RemotePlugin;
}

const getStyles = () => ({
    card: css({
        borderRadius: 16,
        maxWidth: 400,
        margin: 16,
        border: "1px solid #e0e0e0",
    }),
    avatar: css({
        backgroundColor: "#1976d2 !important",
    }),
    chip: css({
        marginRight: 4,
        marginBottom: 4,
    }),
    divider: css({
        margin: "8px 0",
    }),
    orgLink: css({
        color: "inherit",
        textDecoration: "none",
        "&:hover": {
            textDecoration: "underline",
        },
    }),
    actions: css({
        display: "flex",
        justifyContent: "space-between",
    }),
});

export const RemotePluginCard: React.FC<RemotePluginCardProps> = ({
    plugin,
}) => {
    const styles = getStyles();

    const {
        name,
        description,
        orgName,
        orgUrl,
        downloads,
        version,
        verified,
        keywords,
        url,
    } = {
        ...plugin,
        description: "",
        orgName: "",
        orgUrl: "",
        downloads: "",
        version: "",
        verified: "",
        keywords: [],
    };

    return (
        <Card className={styles.card} variant="outlined">
            <CardHeader
                avatar={
                    <Avatar className={styles.avatar}>
                        {name[0].toUpperCase()}
                    </Avatar>
                }
                title={
                    <Stack direction="row" alignItems="center" spacing={1}>
                        <Typography variant="h6">{name}</Typography>
                        {verified && (
                            <Tooltip title="Verified Plugin">
                                <VerifiedIcon
                                    color="primary"
                                    fontSize="small"
                                />
                            </Tooltip>
                        )}
                    </Stack>
                }
                subheader={`v${version}`}
            />

            <CardContent>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                    {description}
                </Typography>

                {keywords.length > 0 && (
                    <Box sx={{ mt: 1, mb: 1 }}>
                        {keywords.map((kw) => (
                            <Chip
                                key={kw}
                                label={kw}
                                size="small"
                                className={styles.chip}
                            />
                        ))}
                    </Box>
                )}

                <Divider className={styles.divider} />

                <Stack direction="row" alignItems="center" spacing={1}>
                    <BusinessIcon fontSize="small" color="action" />
                    <Typography variant="body2" color="text.secondary">
                        <a
                            href={orgUrl}
                            target="_blank"
                            rel="noopener noreferrer"
                            className={styles.orgLink}
                        >
                            {orgName}
                        </a>
                    </Typography>
                </Stack>
            </CardContent>

            <CardActions className={styles.actions}>
                <Stack direction="row" spacing={1} alignItems="center">
                    <DownloadIcon fontSize="small" color="action" />
                    <Typography variant="body2">
                        {downloads.toLocaleString()} downloads
                    </Typography>
                </Stack>
                <Button
                    variant="contained"
                    size="small"
                    endIcon={<LinkIcon />}
                    href={url}
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    View
                </Button>
            </CardActions>
        </Card>
    );
};
