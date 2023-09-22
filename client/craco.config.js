const path = require('path');
module.exports = {
  webpack: {
    alias: {
      '@services': path.resolve(__dirname, 'src/services'),
      "@components/*": path.resolve(__dirname, 'src/components'),
      "@transforms/*": path.resolve(__dirname, "src/transforms"),
      "@mocks/*": path.resolve(__dirname, "src/mocks"),
      "@types/*": path.resolve(__dirname, "src/types"),
    },
  },
};