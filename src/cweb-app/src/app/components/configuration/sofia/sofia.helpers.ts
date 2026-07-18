export function profilesNeedingGatewaySubscription(
  profiles: Record<number, {gateways?: unknown}> | null | undefined,
  requestedProfileIds: ReadonlySet<number>,
): number[] {
  if (!profiles) {
    return [];
  }

  return Object.keys(profiles)
    .map(Number)
    .filter(id => profiles[id].gateways === undefined && !requestedProfileIds.has(id));
}
