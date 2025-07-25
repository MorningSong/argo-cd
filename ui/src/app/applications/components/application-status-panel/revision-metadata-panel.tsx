import {DataLoader, Tooltip} from 'argo-ui';
import * as React from 'react';
import {Timestamp} from '../../../shared/components/timestamp';
import {services} from '../../../shared/services';

export const RevisionMetadataPanel = (props: {appName: string; appNamespace: string; type: string; revision: string; versionId: number}) => {
    if (props.type === 'helm' || props.type === 'oci') {
        return null;
    }
    if (props.type === 'oci') {
        return (
            <DataLoader load={() => services.applications.ociMetadata(props.appName, props.appNamespace, props.revision, 0, props.versionId)} errorRenderer={() => <div />}>
                {m => (
                    <Tooltip
                        popperOptions={{
                            modifiers: {
                                preventOverflow: {
                                    enabled: false
                                },
                                hide: {
                                    enabled: false
                                },
                                flip: {
                                    enabled: false
                                }
                            }
                        }}
                        content={
                            <span>
                                {m.authors && <React.Fragment>Authored by {m.authors}</React.Fragment>}
                                <br />
                                {m.createdAt && <Timestamp date={m.createdAt} />}
                                <br />
                                <br />
                                {m.description}
                            </span>
                        }
                        placement='bottom'
                        allowHTML={true}>
                        <div className='application-status-panel__item-name'>
                            {m.authors && (
                                <div className='application-status-panel__item__row'>
                                    <div>Author:</div>
                                    <div>{m.authors}</div>
                                </div>
                            )}
                        </div>
                    </Tooltip>
                )}
            </DataLoader>
        );
    }
    return (
        <DataLoader
            key={props.revision}
            load={() => services.applications.revisionMetadata(props.appName, props.appNamespace, props.revision, 0, props.versionId)}
            errorRenderer={() => <div />}>
            {m => (
                <Tooltip
                    popperOptions={{
                        modifiers: {
                            preventOverflow: {
                                enabled: false
                            },
                            hide: {
                                enabled: false
                            },
                            flip: {
                                enabled: false
                            }
                        }
                    }}
                    content={
                        <span>
                            {m.author && <React.Fragment>Authored by {m.author}</React.Fragment>}
                            <br />
                            {m.date && <Timestamp date={m.date} />}
                            <br />
                            {m.tags && (
                                <span>
                                    Tags: {m.tags}
                                    <br />
                                </span>
                            )}
                            {m.signatureInfo}
                            <br />
                            {m.message}
                        </span>
                    }
                    placement='bottom'
                    allowHTML={true}>
                    <div className='application-status-panel__item-name'>
                        {m.author && (
                            <div className='application-status-panel__item__row'>
                                <div>Author:</div>
                                <div>
                                    {m.author} - {m.signatureInfo}
                                </div>
                            </div>
                        )}
                        <div className='application-status-panel__item__row'>
                            <div>Comment:</div>
                            <div>{m.message?.split('\n')[0].slice(0, 64)}</div>
                        </div>
                    </div>
                </Tooltip>
            )}
        </DataLoader>
    );
};
