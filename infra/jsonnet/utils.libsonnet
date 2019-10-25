local extractComponent(chal, component) = if component in chal then chal[component] else [];
{
  extractServices(chal):: extractComponent(chal, 'services'),
  extractEmails(chal):: extractComponent(chal, 'emails'),
  extractFlags(chal):: extractComponent(chal, 'flags'),
  extractFiles(chal)::extractComponent(chal, 'staticfiles'),
  extractCLIFiles(chal)::extractComponent(chal, 'clistaticfiles'),
}
