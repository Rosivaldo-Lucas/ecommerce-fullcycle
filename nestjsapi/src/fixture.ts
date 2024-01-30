import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { getDataSourceToken } from '@nestjs/typeorm';
import { DataSource } from 'typeorm';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  await app.init();

  const dataSource = app.get<DataSource>(getDataSourceToken());
  await dataSource.synchronize(true);

  const productRepository = dataSource.getRepository('Product');
  await productRepository.insert([
    {
      id: '0dd4dd1f-268a-4b16-98f8-85897ab44bdf',
      name: 'Product 1',
      description: 'Description 1',
      price: 100,
      image_url: 'http://localhost:9000/products/1.png',
    },
    {
      id: '5890e592-24bb-4222-b601-a76845507dd7',
      name: 'Product 2',
      description: 'Description 2',
      price: 200,
      image_url: 'http://localhost:9000/products/2.png',
    },
  ]);

  await app.close();
}

bootstrap();
