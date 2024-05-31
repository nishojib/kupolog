import styleguide from '@vercel/style-guide/prettier';

const config = { ...styleguide, plugins: [...styleguide.plugins] };

export default config;
