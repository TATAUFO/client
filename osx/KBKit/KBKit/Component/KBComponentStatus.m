//
//  KBComponent.m
//  Keybase
//
//  Created by Gabriel on 5/13/15.
//  Copyright (c) 2015 Gabriel Handford. All rights reserved.
//

#import "KBComponentStatus.h"

#import <ObjectiveSugar/ObjectiveSugar.h>


@interface KBComponentStatus ()
@property NSError *error;
@property KBRInstallStatus installStatus;
@property KBRInstallAction installAction;
@property KBRuntimeStatus runtimeStatus;
@property GHODictionary *info;
@end

@implementation KBComponentStatus

+ (instancetype)componentStatusWithVersion:(KBSemVersion *)version bundleVersion:(KBSemVersion *)bundleVersion runtimeStatus:(KBRuntimeStatus)runtimeStatus info:(GHODictionary *)info {
  if (version && bundleVersion) {
    if ([bundleVersion isGreaterThan:version]) {
      return [KBComponentStatus componentStatusWithInstallStatus:KBRInstallStatusNeedsUpgrade installAction:KBRInstallActionUpgrade runtimeStatus:runtimeStatus info:info error:nil];
    } else {
      return [KBComponentStatus componentStatusWithInstallStatus:KBRInstallStatusInstalled installAction:KBRInstallActionNone runtimeStatus:runtimeStatus info:info error:nil];
    }
  } else if (version && !bundleVersion) {
    return [KBComponentStatus componentStatusWithInstallStatus:KBRInstallStatusInstalled installAction:KBRInstallActionNone runtimeStatus:runtimeStatus info:info error:nil];
  } else if (!version && bundleVersion) {
    return [KBComponentStatus componentStatusWithInstallStatus:KBRInstallStatusNotInstalled installAction:KBRInstallActionInstall runtimeStatus:runtimeStatus info:info error:nil];
  }
  return [KBComponentStatus componentStatusWithInstallStatus:KBRInstallStatusUnknown installAction:KBRInstallActionNone runtimeStatus:runtimeStatus info:info error:nil];
}

+ (instancetype)componentStatusWithInstallStatus:(KBRInstallStatus)installStatus runtimeStatus:(KBRuntimeStatus)runtimeStatus info:(GHODictionary *)info {
  KBRInstallAction installAction;
  switch (installStatus) {
    case KBRInstallStatusError: installAction = KBRInstallActionReinstall; break;
    case KBRInstallStatusUnknown: installAction = KBRInstallActionUnknown; break;
    case KBRInstallStatusInstalled: installAction = KBRInstallActionNone; break;
    case KBRInstallStatusNotInstalled: installAction = KBRInstallActionInstall; break;
    case KBRInstallStatusNeedsUpgrade: installAction = KBRInstallActionUpgrade; break;
  }
  return [self componentStatusWithInstallStatus:installStatus installAction:installAction runtimeStatus:runtimeStatus info:info error:nil];
}

+ (instancetype)componentStatusWithInstallStatus:(KBRInstallStatus)installStatus installAction:(KBRInstallAction)installAction runtimeStatus:(KBRuntimeStatus)runtimeStatus info:(GHODictionary *)info error:(NSError *)error {
  KBComponentStatus *componentStatus = [[KBComponentStatus alloc] init];
  componentStatus.installStatus = installStatus;
  componentStatus.runtimeStatus = runtimeStatus;
  componentStatus.installAction = installAction;
  componentStatus.info = info;
  componentStatus.error = error;
  return componentStatus;
}

+ (instancetype)componentStatusWithServiceStatus:(KBRServiceStatus *)serviceStatus {
  KBComponentStatus *componentStatus = [[KBComponentStatus alloc] init];
  componentStatus.installStatus = serviceStatus.installStatus;
  componentStatus.installAction = serviceStatus.installAction;
  componentStatus.error = serviceStatus.error ? KBMakeError(-1, @"%@", serviceStatus.error.message) : nil;
  componentStatus.runtimeStatus = ![serviceStatus.pid isEqualToString:@""] ? KBRuntimeStatusRunning : KBRuntimeStatusNotRunning;

  GHODictionary *info = [GHODictionary dictionary];
  info[@"Version"] = KBIfBlank(serviceStatus.version, nil);

  if (![serviceStatus.version isEqualToString:serviceStatus.bundleVersion]) {
    info[@"Bundle Version"] = KBIfBlank(serviceStatus.bundleVersion, nil);
  }

  info[@"Label"] = serviceStatus.label;

  //info[@"PID"] = KBIfBlank(serviceStatus.pid, nil);
  info[@"Exit Status"] = KBIfBlank(serviceStatus.lastExitStatus, nil);

  componentStatus.info = info;
  
  return componentStatus;
}

- (BOOL)needsInstallOrUpgrade {
  return _installAction == KBRInstallActionInstall
    || _installAction == KBRInstallActionUpgrade
    || _installAction == KBRInstallActionReinstall;
}

- (GHODictionary *)statusInfo {
  GHODictionary *info = [GHODictionary dictionary];

  if (self.error) {
    info[@"Error"] = [self.error localizedDescription];
  }

  if (_info) [info addEntriesFromOrderedDictionary:_info];
  info[@"Install Status"] = NSStringFromKBRInstallStatus(self.installStatus);
  info[@"Runtime Status"] = NSStringFromKBRuntimeStatus(self.runtimeStatus);
  info[@"Install Action"] = NSStringFromKBRInstallAction(self.installAction);
  return info;
}

- (NSString *)statusDescription {
  NSMutableArray *str = [NSMutableArray array];

  if (_runtimeStatus != KBRuntimeStatusNone) {
    [str addObject:NSStringWithFormat(@"%@, %@", NSStringFromKBRuntimeStatus(_runtimeStatus), NSStringFromKBRInstallStatus(_installStatus))];
  } else {
    [str addObject:NSStringFromKBRInstallStatus(_installStatus)];
  }

  NSMutableArray *infos = [NSMutableArray array];
  for (id key in _info) {
    [infos addObject:NSStringWithFormat(@"%@: %@", key, _info[key])];
  }
  if ([infos count] > 0) [str addObject:NSStringWithFormat(@"\n%@", [infos join:@"\n"])];

  return [str join:@" "];
}

@end

NSString *NSStringFromKBRInstallAction(KBRInstallAction action) {
  switch (action) {
    case KBRInstallActionNone: return nil;
    case KBRInstallActionUnknown: return @"Unknown";
    case KBRInstallActionInstall: return @"Install";
    case KBRInstallActionUpgrade: return @"Upgrade";
    case KBRInstallActionReinstall: return @"Re-Install";
  }
}

NSString *NSStringFromKBRInstallStatus(KBRInstallStatus status) {
  switch (status) {
    case KBRInstallStatusError: return @"Error";
    case KBRInstallStatusUnknown: return @"Unknown";
    case KBRInstallStatusNotInstalled: return @"Not Installed";
    case KBRInstallStatusNeedsUpgrade: return @"Needs Upgrade";
    case KBRInstallStatusInstalled: return @"Installed";
  }
}

NSString *NSStringFromKBRuntimeStatus(KBRuntimeStatus status) {
  switch (status) {
    case KBRuntimeStatusNone: return @"-";
    case KBRuntimeStatusNotRunning: return @"Not Running";
    case KBRuntimeStatusRunning: return @"Running";
  }
}
