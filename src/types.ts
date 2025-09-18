import { PluginType } from "@grafana/data";

export type RemotePlugin = {
    id: number;
    slug: string;
    typeId: number;
    typeName: string;
    typeCode: PluginType;
    name: string;
    url: string;
    description: string;
    orgName: string;
    orgUrl: string;
    keywords?: string[];
    version: string;
    updatedAt: string;
};

// type RemotePlugin = {
//     createdAt: string;
//     description: string;
//     downloads: number;
//     downloadSlug: string;
//     featured: number;
//     id: number;
//     internal: boolean;
//     keywords: string[];
//     json?: {
//         dependencies: PluginDependencies;
//         iam?: IdentityAccessManagement;
//         info: {
//             links: Array<{
//                 name: string;
//                 url: string;
//             }>;
//         };
//     };
//     links: Array<{ rel: string; href: string }>;
//     name: string;
//     orgId: number;
//     orgName: string;
//     orgSlug: string;
//     orgUrl: string;
//     packages: {
//         [arch: string]: {
//             packageName: string;
//             downloadUrl: string;
//         };
//     };
//     popularity: number;
//     readme?: string;
//     signatureType: PluginSignatureType | "";
//     slug: string;
//     status: RemotePluginStatus;
//     statusContext?: string;
//     typeCode: PluginType;
//     typeId: number;
//     typeName: string;
//     updatedAt: string;
//     url: string;
//     userId: number;
//     verified: boolean;
//     version: string;
//     versionSignatureType: PluginSignatureType | "";
//     versionSignedByOrg: string;
//     versionSignedByOrgName: string;
//     versionStatus: string;
//     angularDetected?: boolean;
// };

// interface IdentityAccessManagement {
//     permissions: Permission[];
// }

// interface Permission {
//     action: string;
//     scope: string;
// }

// enum RemotePluginStatus {
//     Deleted = "deleted",
//     Active = "active",
//     Pending = "pending",
//     Deprecated = "deprecated",
//     Enterprise = "enterprise",
// }
